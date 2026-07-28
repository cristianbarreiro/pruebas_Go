package main

import "fmt"

// merge combina dos slices ya ordenados (sortedSlice1, sortedSlice2)
// en un único slice ordenado (mergedSlice).
func merge(sortedSlice1, sortedSlice2, mergedSlice []int) {
	var index1, index2, index3 int

	for index1 < len(sortedSlice1) && index2 < len(sortedSlice2) {
		if sortedSlice1[index1] < sortedSlice2[index2] {
			mergedSlice[index3] = sortedSlice1[index1]
			index1++
			index3++
		} else {
			// caso en que sortedSlice2 tiene el elemento más chico
			mergedSlice[index3] = sortedSlice2[index2]
			index2++
			index3++
		}
	}

	// después de que el for termina, mergeamos el resto del slice
	// que no se depletó (agotó) todavía
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

// mergeSort ordena un slice de forma recursiva: lo divide en dos mitades,
// ordena cada mitad recursivamente (hasta llegar a slices de 1 elemento,
// que ya están "ordenados" por definición), y después las mergea.
func mergeSort(slice []int) []int {
	n := len(slice)

	// caso base: un slice de 0 o 1 elementos ya está ordenado
	if n <= 1 {
		return slice
	}

	mid := n / 2

	// ordenar recursivamente cada mitad
	left := mergeSort(append([]int{}, slice[:mid]...))
	right := mergeSort(append([]int{}, slice[mid:]...))

	// mergear las dos mitades ordenadas
	merged := make([]int, n)
	merge(left, right, merged)

	return merged
}

// mergeSortIterative es la versión iterativa (bottom-up):
// empieza mergeando slices de 1 elemento, después de 2, de 4, etc.
// hasta que todo el slice queda ordenado.
func mergeSortIterative(slice []int) []int {
	n := len(slice)
	result := append([]int{}, slice...)

	for width := 1; width < n; width *= 2 {
		temp := make([]int, n)
		for i := 0; i < n; i += 2 * width {
			mid := i + width
			end := i + 2*width
			if mid > n {
				mid = n
			}
			if end > n {
				end = n
			}
			merge(result[i:mid], result[mid:end], temp[i:end])
		}
		result = temp
	}

	return result
}

func main() {
	datos := []int{9, 4, 7, 1, 3, 8, 2, 6, 5}

	fmt.Println("Slice original:      ", datos)
	fmt.Println("Ordenado (recursivo):", mergeSort(datos))
	fmt.Println("Ordenado (iterativo):", mergeSortIterative(datos))
}
