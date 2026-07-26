package main

import (
	"fmt"
)

// convertToKm convierte millas y yardas a kilómetros
func convertToKm(miles int, yards int) float64 {
	// 1 milla = 1.60934 km
	// 1 yarda = 0.0009144 km
	return float64(miles)*1.60934 + float64(yards)*0.0009144
}

func main() {
	var miles, yards int
	for i := 0; i < 10; i++ {
		fmt.Println("Convert Miles and Yards to Kilometers: Input 2 int")
		fmt.Println("If Miles is negative program ends")

		_, err := fmt.Scanf("%d %d\n", &miles, &yards)
		if err != nil {
			fmt.Println("Invalid input, please enter two integers.")
			continue
		}

		if miles < 0 {
			return
		}

		fmt.Printf("Answer is %.4f kilometers.\n\n", convertToKm(miles, yards))
	}
}
