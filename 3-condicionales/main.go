package main

import "fmt"

func main() {
	edad := 20

	if edad >= 18 {
		fmt.Println("Mayor de edad")
	} else {
		fmt.Println("Menor de edad")
	}

	dia := "lunes"
	switch dia {
	case "lunes":
		fmt.Println("Inicio de semana")
	case "viernes":
		fmt.Println("Fin de semana laboral")
	default:
		fmt.Println("Dia normal")
	}
}
