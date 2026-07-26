package main

import (
	"fmt"
	"strings"
)

func main() {
	var strInput string
	fmt.Println("Palindrome Program Input a string")
	fmt.Scanf("%s", &strInput)
	sum := len(strInput)
	strTest := strings.ToLower(strInput)
	for i := 0; i < sum; i++ {
		fmt.Printf("char %d is %c \n", i, strTest[i]) //prints lower case
		if strTest[i] == strTest[sum-(i+1)] {
			continue
		} else {
			fmt.Println("Not a palindrome")
			return
		}
	}
	fmt.Println("Is a palindrome")
}
