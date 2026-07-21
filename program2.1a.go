package main

import "fmt"

var myInt int = 5 // camel case unlike my_int which is snake case

func main() {
	fmt.Println("Will try some simple ideas")
	fmt.Println("Print the global", myInt)
	{
		var myInt int = 6 //redeclared outer declaration hidden
		fmt.Println("Print the local", myInt)
	}
	fmt.Println("Print the global again", myInt)
}
