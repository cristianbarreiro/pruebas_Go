package main

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func dividir(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "error: division por cero"
	}
	return a / b, "ok"
}

func main() {
	fmt.Println("Suma:", sumar(5, 3))

	resultado, estado := dividir(10, 2)
	fmt.Println("Division:", resultado, "-", estado)

	_, err := dividir(10, 0)
	fmt.Println("Error:", err)
}
