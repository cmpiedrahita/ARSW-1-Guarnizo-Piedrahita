# Laboratorio ARSW - Implementación en Go

## Parte I - Introducción a Hilos en Go

Esta implementación replica la funcionalidad del laboratorio original de Java usando Go y goroutines.

### Archivos implementados:

1. **`threads/count_thread.go`**: Equivalente a `CountThread.java`
   - Define la estructura `CountThread` con campos `start` y `end`
   - Método `Run()` que imprime números en el rango especificado

2. **`main.go`**: Equivalente a `CountThreadsMain.java`
   - Crea 3 hilos con rangos [0..99], [99..199], [200..299]
   - Ejecuta concurrentemente usando goroutines (equivalente a `start()`)

3. **`main_comparison.go`**: Demostración adicional
   - Muestra la diferencia entre ejecución concurrente y secuencial
   - Equivalente a comparar `start()` vs `run()` en Java

### Cómo ejecutar:

```bash
# Ejecución concurrente básica
go run main.go

# Comparación entre concurrente y secuencial
go run main_comparison.go
```

### Diferencias clave entre Go y Java:

- **Goroutines vs Threads**: Go usa goroutines que son más ligeras que los threads de Java
- **sync.WaitGroup**: Equivalente a `join()` en Java para esperar que terminen las goroutines
- **No herencia**: Go no tiene herencia, usa composición y interfaces
- **Concurrencia nativa**: Go tiene concurrencia como característica central del lenguaje

### Observaciones:

- **Ejecución concurrente**: Los números aparecen mezclados porque las goroutines se ejecutan en paralelo
- **Ejecución secuencial**: Los números aparecen en orden porque cada hilo termina antes de que inicie el siguiente

Esto demuestra el mismo comportamiento que se observaría en Java al usar `start()` vs `run()`.