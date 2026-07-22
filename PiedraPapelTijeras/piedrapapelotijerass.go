package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	rock     = 0
	paper    = 1
	scissors = 2
)

func moveName(move int) string {
	switch move {
	case rock:
		return "Rock"
	case paper:
		return "Paper"
	case scissors:
		return "Scissors"
	default:
		return "Unknown"
	}
}

// Devuelve la jugada que gana contra la jugada más frecuente del usuario
func counterMove(userMoves []int) int {
	rockCount := 0
	paperCount := 0
	scissorsCount := 0

	for _, move := range userMoves {
		switch move {
		case rock:
			rockCount++
		case paper:
			paperCount++
		case scissors:
			scissorsCount++
		}
	}

	mostPlayed := rock

	if paperCount > rockCount && paperCount > scissorsCount {
		mostPlayed = paper
	} else if scissorsCount > rockCount && scissorsCount > paperCount {
		mostPlayed = scissors
	}

	// Elegir la jugada que vence a la más usada por el usuario
	switch mostPlayed {
	case rock:
		return paper // papel gana a piedra
	case paper:
		return scissors // tijera gana a papel
	default:
		return rock // piedra gana a tijera
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var userMove int
	var userHistory []int

	var userWins int
	var machineWins int
	var draws int

	for round := 1; round <= 100; round++ {

		fmt.Println("\nRound", round)
		fmt.Println("Choose: 0 = Rock, 1 = Paper, 2 = Scissors")
		fmt.Scan(&userMove)

		if userMove < 0 || userMove > 2 {
			fmt.Println("Invalid move")
			round--
			continue
		}

		userHistory = append(userHistory, userMove)

		var machineMove int

		// Cada 10 rondas la máquina adapta su estrategia
		if round%10 == 0 {
			machineMove = counterMove(userHistory)
			fmt.Println("Machine adapted strategy!")
		} else {
			machineMove = rand.Intn(3)
		}

		fmt.Println("You played:", moveName(userMove))
		fmt.Println("Machine played:", moveName(machineMove))

		if userMove == machineMove {
			fmt.Println("Draw")
			draws++

		} else if (userMove == rock && machineMove == scissors) ||
			(userMove == paper && machineMove == rock) ||
			(userMove == scissors && machineMove == paper) {

			fmt.Println("You win!")
			userWins++

		} else {
			fmt.Println("Machine wins!")
			machineWins++
		}
	}

	fmt.Println("\n===== Final Results =====")
	fmt.Println("User wins:", userWins)
	fmt.Println("Machine wins:", machineWins)
	fmt.Println("Draws:", draws)
}