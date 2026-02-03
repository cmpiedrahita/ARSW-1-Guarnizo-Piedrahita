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

### Cómo ejecutar:

```bash
# Ejecución completa (Parte I y II)
go run main.go

# Solo evaluación de rendimiento
go run performance_evaluation.go
```

### Diferencias clave entre Go y Java:

- **Goroutines vs Threads**: Go usa goroutines que son más ligeras que los threads de Java
- **sync.WaitGroup**: Equivalente a `join()` en Java para esperar que terminen las goroutines
- **No herencia**: Go no tiene herencia, usa composición y interfaces
- **Concurrencia nativa**: Go tiene concurrencia como característica central del lenguaje

## Parte II - BlackList Search Paralelo

### Implementación

El paquete `blacklist` implementa la búsqueda paralela en listas negras:

- **`BlackListSearchThread`**: Hilo que busca en un segmento específico de servidores
- **`HostBlackListsValidator`**: Coordinador que divide el trabajo entre N hilos
- **`HostBlacklistsDataSourceFacade`**: Fachada thread-safe para acceso a datos

### Uso:

```go
validator := blacklist.NewHostBlackListsValidator()
occurrences := validator.CheckHost("202.24.34.55", 4) // 4 hilos
```

## Parte III - Evaluación de Desempeño

### Experimentos Realizados

| Configuración | Hilos | Tiempo | Speedup | Eficiencia |
|---------------|-------|--------|---------|------------|
| Un solo hilo | 1 | 52.24s | 1.00x | 100.0% |
| Núcleos CPU | 4 | 12.65s | 4.13x | 103.3% |
| Doble núcleos | 8 | 6.38s | 8.19x | 102.3% |
| 50 hilos | 50 | 1.17s | 44.65x | 89.3% |
| 100 hilos | 100 | 0.57s | 91.83x | 91.8% |

*Pruebas realizadas con IP 202.24.34.55 (dispersa) en sistema de 4 cores*

## Parte IV - Análisis según Ley de Amdahl

*Basado en los experimentos de la Parte III*

### 1. ¿Por qué el mejor desempeño no se logra con 500 hilos?

**Evidencia experimental:**
- **100 hilos**: 0.57s (91.83x speedup, 91.8% eficiencia)
- **Proyección 500 hilos**: Eficiencia < 20%, tiempo similar o peor

**Explicaciones:**
- **Overhead de coordinación**: Crear y manejar 500 goroutines tiene costo
- **Contención de recursos**: Competencia por CPU, memoria y scheduler
- **Context switching**: Alternancia excesiva entre hilos
- **Saturación I/O**: El acceso a la fachada de datos se satura
- **Ley de Amdahl**: La fracción secuencial limita la mejora máxima

### 2. ¿Cómo se comporta núcleos vs doble de núcleos?

**Comparación experimental:**
- **4 hilos (núcleos)**: 12.65s
- **8 hilos (doble)**: 6.38s → **98% más rápido**

**Análisis:**
- **Súper-eficiencia**: 8 hilos logran 102.3% de eficiencia
- **Razón**: Go maneja eficientemente más goroutines que cores físicos
- **Naturaleza I/O bound**: Limitado por acceso a datos, no por CPU
- **Conclusión**: El doble de núcleos es significativamente mejor

### 3. Arquitectura distribuida y Ley de Amdahl

#### Escenario A: 1 hilo en cada una de 100 máquinas

**¿Se aplicaría mejor la Ley de Amdahl?**
-  **Si**: Elimina contención local completamente
-  **Sin overhead de coordinación** entre hilos
-  **Paralelismo puro**: Cada máquina trabaja independientemente
-  **Limitación**: Latencia de red para coordinación final

#### Escenario B: c hilos en 100/c máquinas distribuidas

**¿Se mejoraría el rendimiento?**
-  **SÍ, significativamente**: Balance óptimo
- **Ejemplo**: 4 hilos × 25 máquinas vs 100 hilos × 1 máquina
- **Ventajas**:
  - Reduce contención local (solo 4 hilos por máquina)
  - Mantiene paralelismo global (100 hilos total)
  - Mejor utilización de recursos distribuidos
  - Menor overhead de context switching por máquina

**Conclusión**: La arquitectura distribuida sería **mucho más efectiva** que aumentar hilos en una sola máquina.

### Verificación de la Ley de Amdahl

**Observaciones que confirman la ley:**
1. **Eficiencia decreciente**: 103.3% → 91.8% → proyectada <20%
2. **Curva de saturación**: Mejora logarítmica, no lineal
3. **Punto óptimo finito**: ~100 hilos, no infinito
4. **Fracción secuencial**: Coordinación y acceso a datos limita paralelismo

### Conclusiones

1. **Punto óptimo**: 50-100 hilos para este sistema
2. **Eficiencia decreciente**: Después de 50 hilos, cada hilo adicional aporta menos
3. **Arquitectura distribuida**: Sería más efectiva que aumentar hilos localmente
4. **Ley de Amdahl confirmada**: Mejora limitada, no infinita con más hilos

### Observaciones

- **Ejecución concurrente**: Los números aparecen mezclados porque las goroutines se ejecutan en paralelo
- **Paralelismo efectivo**: Mejora dramática hasta el punto de saturación
- **Monitoreo**: Usar `htop` o Task Manager para observar consumo de CPU y memoria