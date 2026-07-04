package main

import "fmt"

func main() {
	// Bucle clásico
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteracion:", i)
	}

	// Bucle tipo while
	n := 10
	for n > 0 {
		fmt.Println("Contando:", n)
		n -= 3
	}
}
