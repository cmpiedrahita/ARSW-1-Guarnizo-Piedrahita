// Package blacklist - Implementación del hilo de búsqueda en listas negras
package blacklist

import "sync"

// BlackListSearchThread representa un hilo que busca una IP específica
// en un segmento determinado de servidores de listas negras
type BlackListSearchThread struct {
	startIndex  int   // Índice inicial del segmento de servidores
	endIndex    int   // Índice final del segmento de servidores
	ipAddress   string // Dirección IP a buscar
	occurrences []int  // Lista de servidores donde se encontró la IP
	mu          sync.Mutex // Mutex para acceso thread-safe a occurrences
}

// NewBlackListSearchThread crea una nueva instancia del hilo de búsqueda
// Parámetros:
//   - startIndex: Índice inicial del segmento de servidores a revisar
//   - endIndex: Índice final del segmento de servidores a revisar
//   - ipAddress: Dirección IP a buscar en las listas negras
func NewBlackListSearchThread(startIndex, endIndex int, ipAddress string) *BlackListSearchThread {
	return &BlackListSearchThread{
		startIndex:  startIndex,
		endIndex:    endIndex,
		ipAddress:   ipAddress,
		occurrences: make([]int, 0),
	}
}

// Run ejecuta la búsqueda en el segmento asignado de servidores.
// Itera a través de cada servidor en el rango [startIndex, endIndex)
// y verifica si la IP está en la lista negra de ese servidor.
func (b *BlackListSearchThread) Run() {
	facade := GetInstance()
	
	for i := b.startIndex; i < b.endIndex; i++ {
		if facade.IsInBlackListServer(i, b.ipAddress) {
			// Acceso thread-safe para agregar ocurrencia
			b.mu.Lock()
			b.occurrences = append(b.occurrences, i)
			b.mu.Unlock()
		}
	}
}

// GetOccurrences retorna una copia de la lista de servidores donde se encontró la IP.
// Retorna:
//   - []int: Copia de los números de servidores donde se encontró la IP
func (b *BlackListSearchThread) GetOccurrences() []int {
	b.mu.Lock()
	defer b.mu.Unlock()
	result := make([]int, len(b.occurrences))
	copy(result, b.occurrences)
	return result
}

// GetOccurrencesCount retorna el número total de ocurrencias encontradas.
// Retorna:
//   - int: Número de servidores donde se encontró la IP
func (b *BlackListSearchThread) GetOccurrencesCount() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return len(b.occurrences)
}