package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var move, machineMove int

	const (
		rock     = 0
		paper    = 1
		scissors = 2
	)

	const (
		cRock     = "R"
		cPaper    = "P"
		cScissors = "S"
	)

	var cMove string
	var draws, wins, machineWins int
	var rounds int

	fmt.Print("How many rounds do you want to play? ")
	fmt.Scan(&rounds)

	// Inicializar números aleatorios
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < rounds; i++ {

		// User plays
		fmt.Println("\nRound", i+1, ": Choose either R, P or S")
		fmt.Scan(&cMove)

		if cMove == cRock {
			move = rock
		} else if cMove == cPaper {
			move = paper
		} else if cMove == cScissors {
			move = scissors
		} else {
			fmt.Println("-> illegal move")
			i--
			continue
		}

		// Machine plays
		machineMove = rand.Intn(3)

		if machineMove == rock {
			cMove = cRock
		} else if machineMove == paper {
			cMove = cPaper
		} else {
			cMove = cScissors
		}

		fmt.Printf("Machine plays %s\n", cMove)

		// Determine result
		if move == machineMove {
			fmt.Println("-> Draw")
			draws++
		} else if (move == paper && machineMove == scissors) ||
			(move == rock && machineMove == paper) ||
			(move == scissors && machineMove == rock) {

			machineWins++
			fmt.Println("-> Machine wins")

		} else {
			wins++
			fmt.Println("-> You win")
		}
	}

	fmt.Println("\nAfter", rounds, "rounds, you won", wins,
		"times, machine won", machineWins,
		"times and there were", draws, "draws.")
}
