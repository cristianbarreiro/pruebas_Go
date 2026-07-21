package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mismoCumpleanos(personas int) bool {
	cumpleanos := make(map[int]bool)

	for i := 0; i < personas; i++ {
		dia := rand.Intn(365) + 1

		if cumpleanos[dia] {
			return true
		}

		cumpleanos[dia] = true
	}

	return false
}

func calcularProbabilidad(personas int, simulaciones int) float64 {
	aciertos := 0

	for i := 0; i < simulaciones; i++ {
		if mismoCumpleanos(personas) {
			aciertos++
		}
	}

	return float64(aciertos) / float64(simulaciones)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const simulaciones = 10000

	fmt.Println("Personas en la sala\tProbabilidad de que 2 (o más) cumplan años el mismo día")

	for personas := 10; personas <= 100; personas += 10 {
		probabilidad := calcularProbabilidad(personas, simulaciones)

		fmt.Printf("%d\t\t\t%.3f\n", personas, probabilidad)
	}
}
