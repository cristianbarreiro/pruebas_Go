// Programa que implementa el algoritmo Merge Sort para ordenar
// una lista de 10.000 números enteros generados pseudoaleatoriamente.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// merge combina dos partes ordenadas de un slice en un único slice ordenado.
func merge(left []int, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// mergeSort ordena un slice utilizando el algoritmo Merge Sort.
func mergeSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	middle := len(numbers) / 2

	left := mergeSort(numbers[:middle])
	right := mergeSort(numbers[middle:])

	return merge(left, right)
}

func main() {

	// Inicializa el generador de números pseudoaleatorios.
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Crea un slice con 10.000 números enteros aleatorios.
	numbers := make([]int, 10000)

	for i := range numbers {
		numbers[i] = random.Intn(100000)
	}

	fmt.Println("Primeros 20 elementos sin ordenar:")
	fmt.Println(numbers[:20])

	// Ordena el slice utilizando Merge Sort.
	sortedNumbers := mergeSort(numbers)

	fmt.Println("\nPrimeros 20 elementos ordenados:")
	fmt.Println(sortedNumbers[:20])
}