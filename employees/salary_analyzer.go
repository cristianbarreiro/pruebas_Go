/*
 Introductory Comment:
 This program reads employee data (first name, last name, salary) from a 
 text file named "employees.txt". It parses this data and stores it in a 
 map where the key is the concatenated full name and the value is the 
 integer salary.
 
 After populating the map, it iterates through the records to find the 
 employee with the lowest salary, the employee with the highest salary, 
 and calculates the company's overall average salary.
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 1. Open the file containing the employee data
	file, err := os.Open("employees.txt")
	if err != nil {
		fmt.Println("Error opening file. Please ensure 'employees.txt' is in the same directory.")
		fmt.Println("Error details:", err)
		return
	}
	defer file.Close()

	// 2. Initialize a map to store the data (Key: Full Name string, Value: Salary int)
	employeeSalaries := make(map[string]int)

	// 3. Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		
		// Split the line by whitespace
		fields := strings.Fields(line)
		
		// Ensure the line has exactly 3 fields before processing
		if len(fields) == 3 {
			firstName := fields[0]
			lastName := fields[1]
			
			// Convert the salary string to an integer
			salary, err := strconv.Atoi(fields[2])
			
			if err == nil {
				// Concatenate name for the map key
				fullName := firstName + " " + lastName
				employeeSalaries[fullName] = salary
			}
		}
	}

	// Catch any errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Exit if the map is empty to avoid division by zero later
	if len(employeeSalaries) == 0 {
		fmt.Println("No valid employee data found in the file.")
		return
	}

	// 4. Variables to track the statistics
	// Set minSalary to the maximum possible 32-bit int so any real salary will be lower
	minSalary := math.MaxInt32 
	maxSalary := -1
	minName := ""
	maxName := ""
	totalSalary := 0

	// 5. Iterate through the map to calculate min, max, and total
	for name, salary := range employeeSalaries {
		totalSalary += salary

		if salary > maxSalary {
			maxSalary = salary
			maxName = name
		}
		
		if salary < minSalary {
			minSalary = salary
			minName = name
		}
	}

	// 6. Calculate the average
	averageSalary := float64(totalSalary) / float64(len(employeeSalaries))

	// 7. Output the required results
	fmt.Println("=========================================")
	fmt.Println("         COMPANY SALARY REPORT           ")
	fmt.Println("=========================================")
	fmt.Printf("Total Employees Evaluated: %d\n\n", len(employeeSalaries))
	fmt.Printf("Lowest Salary:  %s ($%d)\n", minName, minSalary)
	fmt.Printf("Highest Salary: %s ($%d)\n", maxName, maxSalary)
	fmt.Printf("Average Salary: $%.2f\n", averageSalary)
	fmt.Println("=========================================")
}