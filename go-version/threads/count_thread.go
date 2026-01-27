package threads

import "fmt"

// CountThread representa un hilo que imprime números en un rango
type CountThread struct {
	start int
	end   int
}

// NewCountThread crea una nueva instancia de CountThread
func NewCountThread(start, end int) *CountThread {
	return &CountThread{
		start: start,
		end:   end,
	}
}

// Run ejecuta el conteo de números en el rango especificado
func (ct *CountThread) Run() {
	for i := ct.start; i <= ct.end; i++ {
		fmt.Println(i)
	}
}