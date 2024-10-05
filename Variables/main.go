package main

import (
	"fmt"
)

func CalcularPuntajeTotal(puntos[] int) int {
	total := 0
	for _, punto := range puntos {
		total += punto
	}
	return total
}

func ClasificarRendimiento(PuntajeTotal int) string {
	if PuntajeTotal >= 200 {
		return "Excelente"
	} else if PuntajeTotal >= 100 {
		return "Bueno"
	} else if PuntajeTotal >= 50 {
		return "Regular"
	} else {
		return "Maloaaaa"
	}
}

func main() {
	var numRondas int

	fmt.Print("Cuantas rondas jugaste?: ")
	fmt.Scan(&numRondas)

	puntos := make([]int, numRondas)

	for i := 0; i < numRondas; i++ {
		fmt.Printf("Ingresa los puntos de la ronda %d: ", i+1)
		fmt.Scan(&puntos[i])
	}

	PuntajeTotal := CalcularPuntajeTotal(puntos)

	rendimiento := ClasificarRendimiento(PuntajeTotal)

	fmt.Printf("El puntaje total acumulado es: %d\n", PuntajeTotal)
	fmt.Printf("El rendimiento es: %s\n", rendimiento)
}