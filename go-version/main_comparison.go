package main

import (
	"blacklist-search/threads"
	"fmt"
	"sync"
)

func main() {
	// Crear 3 hilos con los rangos especificados
	thread1 := threads.NewCountThread(0, 99)
	thread2 := threads.NewCountThread(99, 199)
	thread3 := threads.NewCountThread(200, 299)

	fmt.Println("=== Ejecución CONCURRENTE (equivalente a start()) ===")
	runConcurrent(thread1, thread2, thread3)

	fmt.Println("\n=== Ejecución SECUENCIAL (equivalente a run()) ===")
	runSequential(thread1, thread2, thread3)
}

func runConcurrent(countThreads ...*threads.CountThread) {
	var wg sync.WaitGroup
	wg.Add(len(countThreads))

	for _, thread := range countThreads {
		go func(t *threads.CountThread) {
			defer wg.Done()
			t.Run()
		}(thread)
	}

	wg.Wait()
}

func runSequential(countThreads ...*threads.CountThread) {
	for _, thread := range countThreads {
		thread.Run()
	}
}