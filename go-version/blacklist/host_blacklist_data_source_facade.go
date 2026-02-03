// Package blacklist - Fachada para acceso a fuentes de datos de listas negras
// Esta implementación simula el comportamiento de la clase Java original
// HostBlacklistsDataSourceFacade, proporcionando acceso thread-safe a
// múltiples listas negras de servidores maliciosos.
package blacklist

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// HostBlacklistsDataSourceFacade proporciona una interfaz unificada
// para consultar múltiples listas negras de hosts maliciosos.
// Implementa el patrón Singleton y es thread-safe.
type HostBlacklistsDataSourceFacade struct {
	blacklistOccurrences map[string]bool // Mapa de ocurrencias: "servidor-IP" -> existe
	mu                   sync.RWMutex    // Mutex para acceso concurrente seguro
}

var (
	instance *HostBlacklistsDataSourceFacade // Instancia singleton
	once     sync.Once                       // Garantiza inicialización única
)

// GetInstance retorna la instancia singleton de la fachada.
// Utiliza sync.Once para garantizar inicialización thread-safe.
func GetInstance() *HostBlacklistsDataSourceFacade {
	once.Do(func() {
		instance = &HostBlacklistsDataSourceFacade{
			blacklistOccurrences: make(map[string]bool),
		}
		instance.initializeBlacklists()
	})
	return instance
}

// initializeBlacklists inicializa las listas negras con datos de prueba.
// Simula el comportamiento del código Java original con IPs específicas
// distribuidas en diferentes servidores para probar el paralelismo.
func (h *HostBlacklistsDataSourceFacade) initializeBlacklists() {
	// IPs encontradas rápidamente (primeros servidores) - 200.24.34.55
	h.blacklistOccurrences["23-200.24.34.55"] = true
	h.blacklistOccurrences["50-200.24.34.55"] = true
	h.blacklistOccurrences["200-200.24.34.55"] = true
	h.blacklistOccurrences["1000-200.24.34.55"] = true
	h.blacklistOccurrences["500-200.24.34.55"] = true
	
	// IPs dispersas en el rango completo - 202.24.34.55
	h.blacklistOccurrences["29-202.24.34.55"] = true
	h.blacklistOccurrences["10034-202.24.34.55"] = true
	h.blacklistOccurrences["20200-202.24.34.55"] = true
	h.blacklistOccurrences["31000-202.24.34.55"] = true
	h.blacklistOccurrences["70500-202.24.34.55"] = true
	// Nota: 212.24.24.55 no está en ninguna lista (caso de prueba)
}

// GetRegisteredServersCount retorna el número total de servidores
// de listas negras disponibles para consulta.
func (h *HostBlacklistsDataSourceFacade) GetRegisteredServersCount() int {
	return 80000
}

// IsInBlackListServer verifica si una IP está registrada en un servidor
// específico de lista negra.
// Parámetros:
//   - serverNumber: Número del servidor de lista negra a consultar
//   - ip: Dirección IP a verificar
// Retorna:
//   - bool: true si la IP está en la lista negra del servidor
func (h *HostBlacklistsDataSourceFacade) IsInBlackListServer(serverNumber int, ip string) bool {
	// Simular latencia de red mínima
	time.Sleep(time.Nanosecond)
	
	key := fmt.Sprintf("%d-%s", serverNumber, ip)
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.blacklistOccurrences[key]
}

// ReportAsNotTrustworthy registra un host como no confiable.
// Este método simula el reporte a una base de datos local.
func (h *HostBlacklistsDataSourceFacade) ReportAsNotTrustworthy(host string) {
	log.Printf("HOST %s Reported as NOT trustworthy", host)
}

// ReportAsTrustworthy registra un host como confiable.
// Este método simula el reporte a una base de datos local.
func (h *HostBlacklistsDataSourceFacade) ReportAsTrustworthy(host string) {
	log.Printf("HOST %s Reported as trustworthy", host)
}