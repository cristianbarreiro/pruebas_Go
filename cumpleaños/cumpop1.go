// Birthday Probability Simulation uses random number generation to estimate
// the probability that at least two people in a room share the same birthday.
// The simulation tests room sizes from 10 to 100 people and performs
// 10,000 trials for each room size.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	daysInYear        = 365
	simulationCount   = 10000
	minimumRoomSize   = 10
	maximumRoomSize   = 100
	roomSizeIncrement = 10
)

func main() {
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Printf(
		"%-22s %s\n",
		"People in the room",
		"Probability of 2 or more having the same birthday",
	)
	fmt.Println("--------------------------------------------------------------------------")

	for numberOfPeople := minimumRoomSize; numberOfPeople <= maximumRoomSize; numberOfPeople += roomSizeIncrement {

		probability := estimateBirthdayProbability(
			numberOfPeople,
			simulationCount,
			randomGenerator,
		)

		fmt.Printf("%-22d %.4f\n", numberOfPeople, probability)
	}
}

// estimateBirthdayProbability estimates the probability that at least two
// people in a room have the same birthday.
func estimateBirthdayProbability(
	numberOfPeople int,
	numberOfSimulations int,
	randomGenerator *rand.Rand,
) float64 {
	simulationsWithMatch := 0

	for simulation := 0; simulation < numberOfSimulations; simulation++ {
		if hasSharedBirthday(numberOfPeople, randomGenerator) {
			simulationsWithMatch++
		}
	}

	return float64(simulationsWithMatch) / float64(numberOfSimulations)
}

// hasSharedBirthday returns true when at least two people share a birthday.
func hasSharedBirthday(
	numberOfPeople int,
	randomGenerator *rand.Rand,
) bool {
	birthdayUsed := make([]bool, daysInYear)

	for person := 0; person < numberOfPeople; person++ {
		birthday := randomGenerator.Intn(daysInYear)

		if birthdayUsed[birthday] {
			return true
		}

		birthdayUsed[birthday] = true
	}

	return false
}
