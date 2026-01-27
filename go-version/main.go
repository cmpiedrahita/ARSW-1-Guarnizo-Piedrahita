package main

import (
	"blacklist-search/threads"
	"sync"
)

func main() {
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