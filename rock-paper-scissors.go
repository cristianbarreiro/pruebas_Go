package main

import "fmt"

func main() {

	var move rune
	fmt.Println("Choose either R P or S")
	fmt.Scanf("%c", &move)
	if move == 'R' {
		fmt.Printf("my move is %c\n", move)
	} else if move == 'P' {
		fmt.Printf("my move is %c\n", move)
	} else if move == 'S' {
		fmt.Printf("my move is %c\n", move)
	} else {
		fmt.Printf("illegal move id %c\n", move)
	}

}
