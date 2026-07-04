package main

import "fmt"

func main() {
	// Array fijo
	colores := [3]string{"rojo", "verde", "azul"}
	fmt.Println("Array:", colores)

	// Slice dinámico
	numeros := []int{1, 2, 3}
	numeros = append(numeros, 4, 5)
	fmt.Println("Slice:", numeros)

	// Recorrer con range
	for i, v := range numeros {
		fmt.Printf("numeros[%d] = %d\n", i, v)
	}
}
