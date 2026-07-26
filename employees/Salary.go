//This class requires a file called emoloyees.txt, please download it from coursera course

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Main function
func main() {
	fileReadResult := ReadFromFile("employees.txt")
	highestKey, highest, lowerKey, lower, total := GetLowestHighestTotal(fileReadResult)
	prom := total / len(fileReadResult)
	fmt.Printf("Highest Key: %s %.d\n\n", highestKey, highest)
	fmt.Printf("Lowest Key: %s %.d\n\n", lowerKey, lower)
	fmt.Printf("Prom: %.d\n\n", prom)
}

// ReadFromFile Read file and generate the map based on the length of the file, return map
func ReadFromFile(filename string) map[string]int {
	result := make(map[string]int)
	//Opens the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return result
	}
	defer file.Close()

	//Reading File
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		if len(parts) == 3 {
			key := strings.TrimSpace(parts[0] + " " + parts[1])
			value, err := strconv.ParseFloat(parts[2], 32)
			if err != nil {
				fmt.Printf("Error Parsing Salary: %s\n", err)
			}
			result[key] = int(value)
		}
	}
	return result
}

// GetLowestHighestTotal Input: Map[string]int
// Returns: string (highestKey), int(highest), string(lowerKey), int(lower), int(total)
func GetLowestHighestTotal(input map[string]int) (string, int, string, int, int) {
	var highest, lower, total int
	var highestKey, lowerKey string
	first := true
	for key, value := range input {
		total += value
		if first {
			highest = value
			lower = value
			highestKey = key
			lowerKey = key
			first = false
			continue
		}
		if value > highest {
			highest = value
			highestKey = key
		}

		if value < lower {
			lower = value
			lowerKey = key
		}
	}
	return highestKey, highest, lowerKey, lower, total
}
