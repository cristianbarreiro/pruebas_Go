// Employee Salary Analyzer reads employee names and salaries from a file.
// It stores the data in a map and prints the employees with the lowest
// and highest salaries, along with the company's average salary.

package main

import (
	"bufio"
	"fmt"
	"os"
)

const employeeFileName = "employees"

func main() {
	employeeSalaries, err := readEmployeeSalaries(employeeFileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(employeeSalaries) == 0 {
		fmt.Println("The employee file contains no data.")
		return
	}

	lowestPaidEmployee, lowestSalary,
		highestPaidEmployee, highestSalary,
		averageSalary := analyzeSalaries(employeeSalaries)

	fmt.Printf("Number of employees: %d\n", len(employeeSalaries))
	fmt.Printf(
		"Employee with the smallest salary: %s ($%d)\n",
		lowestPaidEmployee,
		lowestSalary,
	)
	fmt.Printf(
		"Employee with the largest salary: %s ($%d)\n",
		highestPaidEmployee,
		highestSalary,
	)
	fmt.Printf("Company's average salary: $%.2f\n", averageSalary)
}

// readEmployeeSalaries reads employee data from the specified file.
// Each line must contain a first name, last name, and integer salary.
func readEmployeeSalaries(fileName string) (map[string]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not open %q: %w", fileName, err)
	}
	defer file.Close()

	employeeSalaries := make(map[string]int)

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++

		var firstName string
		var lastName string
		var salary int

		itemsRead, err := fmt.Sscanf(
			scanner.Text(),
			"%s %s %d",
			&firstName,
			&lastName,
			&salary,
		)
		if err != nil || itemsRead != 3 {
			return nil, fmt.Errorf(
				"invalid employee data on line %d: %q",
				lineNumber,
				scanner.Text(),
			)
		}

		employeeName := firstName + "." + lastName
		employeeSalaries[employeeName] = salary
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("could not read %q: %w", fileName, err)
	}

	return employeeSalaries, nil
}

// analyzeSalaries finds the employees with the smallest and largest salaries
// and calculates the average salary for all employees.
func analyzeSalaries(
	employeeSalaries map[string]int,
) (
	lowestPaidEmployee string,
	lowestSalary int,
	highestPaidEmployee string,
	highestSalary int,
	averageSalary float64,
) {
	firstEmployee := true
	totalSalary := 0

	for employeeName, salary := range employeeSalaries {
		if firstEmployee {
			lowestPaidEmployee = employeeName
			highestPaidEmployee = employeeName
			lowestSalary = salary
			highestSalary = salary
			firstEmployee = false
		}

		if salary < lowestSalary {
			lowestSalary = salary
			lowestPaidEmployee = employeeName
		}

		if salary > highestSalary {
			highestSalary = salary
			highestPaidEmployee = employeeName
		}

		totalSalary += salary
	}

	averageSalary = float64(totalSalary) / float64(len(employeeSalaries))

	return
}
