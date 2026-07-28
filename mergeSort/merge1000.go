package main

import (
	"fmt"
	"math/rand"
	"time"
)

// merge combina dos slices ordenados en un único slice ordenado.
func merge(sortedSlice1, sortedSlice2, mergedSlice []int) {
	var index1, index2, index3 int

	for index1 < len(sortedSlice1) && index2 < len(sortedSlice2) {
		if sortedSlice1[index1] < sortedSlice2[index2] {
			mergedSlice[index3] = sortedSlice1[index1]
			index1++
		} else {
			mergedSlice[index3] = sortedSlice2[index2]
			index2++
		}
		index3++
	}

	for index1 < len(sortedSlice1) {
		mergedSlice[index3] = sortedSlice1[index1]
		index1++
		index3++
	}

	for index2 < len(sortedSlice2) {
		mergedSlice[index3] = sortedSlice2[index2]
		index2++
		index3++
	}
}

// mergeSort ordena un slice utilizando el algoritmo Merge Sort.
func mergeSort(slice []int) []int {

	if len(slice) <= 1 {
		return slice
	}

	mid := len(slice) / 2

	left := mergeSort(slice[:mid])
	right := mergeSort(slice[mid:])

	merged := make([]int, len(slice))

	merge(left, right, merged)

	return merged
}

func main() {

	// Genera números pseudoaleatorios.
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Crea un slice desordenado de 10.000 elementos.
	datos := make([]int, 10000)

	for i := range datos {
		datos[i] = random.Intn(100000)
	}

	fmt.Println("Primeros 20 elementos sin ordenar:")
	fmt.Println(datos[:20])

	// Ordena usando Merge Sort.
	datosOrdenados := mergeSort(datos)

	fmt.Println("\nPrimeros 20 elementos ordenados:")
	fmt.Println(datosOrdenados[:20])
}
