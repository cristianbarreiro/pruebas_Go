package main

import "fmt"

func main() {
	capitales := map[string]string{
		"Mexico":    "Ciudad de Mexico",
		"Argentina": "Buenos Aires",
		"España":    "Madrid",
	}

	fmt.Println("Capitales:", capitales)
	fmt.Println("Capital de Mexico:", capitales["Mexico"])

	capitales["Colombia"] = "Bogota"
	delete(capitales, "España")

	for pais, capital := range capitales {
		fmt.Printf("%s -> %s\n", pais, capital)
	}
}
