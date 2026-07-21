/*
 Introductory Comment:
 This program uses a Monte Carlo simulation to estimate the probability
 that at least two people in a room share the same birthday (commonly
 known as the Birthday Paradox).

 It simulates room sizes from 10 to 100 in increments of 10. For each
 room size, it runs 10,000 trials to compute a highly accurate probability.
 The random number generator is properly seeded with the current time
 (UnixNano) to ensure different, truly randomized results on every run.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator to ensure different results per run
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	const numTrials = 10000
	const daysInYear = 365

	// Print the table header exactly as requested
	fmt.Printf("%-20s %s\n", "People in the room", "Probability of 2 (or more) having the same birthday")

	// Loop through room sizes from 10 to 100 in increments of 10
	for peopleCount := 10; peopleCount <= 100; peopleCount += 10 {
		matchCount := 0

		// Run 10,000 simulations for the current room size
		for trial := 0; trial < numTrials; trial++ {
			if hasSharedBirthday(peopleCount, daysInYear, rng) {
				matchCount++
			}
		}

		// Calculate and print the probability (0.xyz format)
		probability := float64(matchCount) / float64(numTrials)
		fmt.Printf("%-20d %.3f\n", peopleCount, probability)
	}
}

// hasSharedBirthday simulates a room of 'peopleCount' and returns true
// if at least two people share a birthday.
func hasSharedBirthday(peopleCount int, daysInYear int, rng *rand.Rand) bool {
	// A boolean slice to keep track of birthdays we have already seen in this trial.
	// We use daysInYear + 1 so we can 1-index the days (1 to 365).
	seen := make([]bool, daysInYear+1)

	for i := 0; i < peopleCount; i++ {
		// Generate a random birthday between 1 and 365
		birthday := rng.Intn(daysInYear) + 1

		// If this birthday has already been marked 'true', we have a match
		if seen[birthday] {
			return true
		}

		// Otherwise, mark this birthday as seen and continue
		seen[birthday] = true
	}

	// No shared birthdays found in this room
	return false
}
