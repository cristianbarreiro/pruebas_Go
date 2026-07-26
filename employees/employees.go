package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Abrir archivo con los salarios
	file, err := os.Open("employees.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	// Crear mapa: nombre -> salario
	salaries := make(map[string]int)

	// Leer archivo línea por línea
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()

		// Separar nombre, apellido y salario
		data := strings.Fields(line)

		if len(data) != 3 {
			fmt.Println("Invalid line:", line)
			continue
		}

		name := data[0] + "." + data[1]

		salary, err := strconv.Atoi(data[2])

		if err != nil {
			fmt.Println("Invalid salary")
			return
		}

		// Guardar en el mapa
		salaries[name] = salary
	}


	// Variables para encontrar mínimo y máximo

	var lowestName string
	var highestName string

	lowestSalary := 999999999
	highestSalary := 0

	totalSalary := 0


	// Recorrer mapa

	for name, salary := range salaries {

		totalSalary += salary

		if salary < lowestSalary {
			lowestSalary = salary
			lowestName = name
		}

		if salary > highestSalary {
			highestSalary = salary
			highestName = name
		}
	}


	// Calcular promedio

	if len(salaries) == 0 {
		fmt.Println("No employees found")
		return
	}

	average := totalSalary / len(salaries)


	// Mostrar resultados

	fmt.Println("Lowest salary:")
	fmt.Println(lowestName, lowestSalary)

	fmt.Println()

	fmt.Println("Highest salary:")
	fmt.Println(highestName, highestSalary)

	fmt.Println()

	fmt.Println("Average salary:")
	fmt.Println(average)

}