package main

import (
	"errors"
	"fmt"
)

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("no se puede dividir por cero")
	}
	return a / b, nil
}

func main() {
	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}

	_, err = dividir(5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
