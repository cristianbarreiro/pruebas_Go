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

	// Inicializar números aleatorios una sola vez
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
			fmt.Println("Machine plays R")
		} else if machineMove == paper {
			fmt.Println("Machine plays P")
		} else {
			fmt.Println("Machine plays S")
		}

		// Determine result using switch
		switch move {

		case rock:
			if machineMove == rock {
				fmt.Println("-> draw")
				draws++
			} else if machineMove == paper {
				fmt.Println("-> machine wins")
				machineWins++
			} else {
				fmt.Println("-> you win")
				wins++
			}

		case paper:
			if machineMove == paper {
				fmt.Println("-> draw")
				draws++
			} else if machineMove == scissors {
				fmt.Println("-> machine wins")
				machineWins++
			} else {
				fmt.Println("-> you win")
				wins++
			}

		case scissors:
			if machineMove == scissors {
				fmt.Println("-> draw")
				draws++
			} else if machineMove == rock {
				fmt.Println("-> machine wins")
				machineWins++
			} else {
				fmt.Println("-> you win")
				wins++
			}
		}
	}

	fmt.Println("\nAfter", rounds, "rounds:")
	fmt.Println("you won", wins, "times")
	fmt.Println("you drew", draws, "times")
	fmt.Println("machine won", machineWins, "times")
}
