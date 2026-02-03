// Package blacklist implementa la funcionalidad de validación de hosts
// contra listas negras de servidores maliciosos usando paralelismo.
package blacklist

import (
	"log"
	"sync"
)

// BLACK_LIST_ALARM_COUNT define el número mínimo de ocurrencias
// para considerar un host como no confiable
const BLACK_LIST_ALARM_COUNT = 5

// HostBlackListsValidator valida direcciones IP contra múltiples listas negras
// utilizando búsqueda paralela para mejorar el rendimiento
type HostBlackListsValidator struct{}

// NewHostBlackListsValidator crea una nueva instancia del validador
func NewHostBlackListsValidator() *HostBlackListsValidator {
	return &HostBlackListsValidator{}
}

// CheckHost valida una dirección IP contra las listas negras usando N hilos paralelos.
// Parámetros:
//   - ipAddress: La dirección IP a validar
//   - numThreads: Número de hilos para paralelizar la búsqueda
// Retorna:
//   - []int: Lista de números de servidores donde se encontró la IP
func (h *HostBlackListsValidator) CheckHost(ipAddress string, numThreads int) []int {
	facade := GetInstance()
	totalServers := facade.GetRegisteredServersCount()
	
	// Calcular el tamaño de segmento para cada hilo
	segmentSize := totalServers / numThreads
	remainder := totalServers % numThreads
	
	threads := make([]*BlackListSearchThread, numThreads)
	var wg sync.WaitGroup
	
	// Crear y ejecutar hilos para búsqueda paralela
	for i := 0; i < numThreads; i++ {
		startIndex := i * segmentSize
		endIndex := startIndex + segmentSize
		// El último hilo maneja los servidores restantes
		if i == numThreads-1 {
			endIndex += remainder
		}
		
		threads[i] = NewBlackListSearchThread(startIndex, endIndex, ipAddress)
		wg.Add(1)
		
		// Ejecutar búsqueda en goroutine separada
		go func(thread *BlackListSearchThread) {
			defer wg.Done()
			thread.Run()
		}(threads[i])
	}
	
	// Esperar a que todos los hilos terminen
	wg.Wait()
	
	// Recopilar resultados de todos los hilos
	var allOccurrences []int
	totalOccurrences := 0
	checkedListsCount := 0
	
	for _, thread := range threads {
		occurrences := thread.GetOccurrences()
		allOccurrences = append(allOccurrences, occurrences...)
		totalOccurrences += len(occurrences)
		checkedListsCount += thread.endIndex - thread.startIndex
	}
	
	// Reportar resultado basado en el número de ocurrencias encontradas
	if totalOccurrences >= BLACK_LIST_ALARM_COUNT {
		facade.ReportAsNotTrustworthy(ipAddress)
	} else {
		facade.ReportAsTrustworthy(ipAddress)
	}
	
	// Log de estadísticas de búsqueda
	log.Printf("Checked Black Lists: %d of %d", checkedListsCount, totalServers)
	
	return allOccurrences
}