package main

import (
	"blacklist-search/blacklist"
	"blacklist-search/threads"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== PARTE I - Hilos Básicos ===")
	runPartI()
	
	fmt.Println("\n=== PARTE II - BlackList Search ===")
	runPartII()
}

func runPartI() {
	// Crear 3 hilos con los rangos especificados
	thread1 := threads.NewCountThread(0, 99)
	thread2 := threads.NewCountThread(99, 199)
	thread3 := threads.NewCountThread(200, 299)

	// WaitGroup para esperar que todas las goroutines terminen
	var wg sync.WaitGroup
	wg.Add(3)

	// Ejecutar con goroutines (equivalente a start() en Java)
	go func() {
		defer wg.Done()
		thread1.Run()
	}()

	go func() {
		defer wg.Done()
		thread2.Run()
	}()

	go func() {
		defer wg.Done()
		thread3.Run()
	}()

	// Esperar que todas las goroutines terminen
	wg.Wait()
}

func runPartII() {
	validator := blacklist.NewHostBlackListsValidator()
	
	// Prueba básica
	fmt.Println("Testing basic functionality:")
	occurrences := validator.CheckHost("200.24.34.55", 4)
	fmt.Printf("Host found in blacklists: %v\n\n", occurrences)
	
	// Prueba de rendimiento
	fmt.Printf("Performance test with dispersed IP (202.24.34.55):\n")
	fmt.Printf("CPU cores: %d\n\n", runtime.NumCPU())
	
	threadCounts := []int{1, runtime.NumCPU(), runtime.NumCPU() * 2}
	
	for _, numThreads := range threadCounts {
		fmt.Printf("Testing with %d threads: ", numThreads)
		start := time.Now()
		occurrences := validator.CheckHost("202.24.34.55", numThreads)
		duration := time.Since(start)
		fmt.Printf("%v (found: %d)\n", duration, len(occurrences))
	}
}