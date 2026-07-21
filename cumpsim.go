package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	daysInYear     = 365
	numberOfTrials = 10000
)

// hasDuplicateBirthday returns true if two or more people
// share the same birthday.
func hasDuplicateBirthday(people int) bool {
	birthdays := make(map[int]bool)

	for i := 0; i < people; i++ {
		// Birthday is a number from 1 to 365
		day := rand.Intn(daysInYear) + 1

		if birthdays[day] {
			return true
		}
		birthdays[day] = true
	}
	return false
}

func main() {
	// Seed the random number generator for different results per run
	rand.Seed(time.Now().UnixNano())

	fmt.Println("People in the room\tProbability of shared birthday")

	for people := 10; people <= 100; people += 10 {
		matches := 0

		for trial := 0; trial < numberOfTrials; trial++ {
			if hasDuplicateBirthday(people) {
				matches++
			}
		}

		probability := float64(matches) / float64(numberOfTrials)
		fmt.Printf("%d\t\t\t%.3f\n", people, probability)
	}
}
