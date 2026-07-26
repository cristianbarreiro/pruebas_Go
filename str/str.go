package main

import (
	"fmt"
)

func main() {
	var c byte = 'a'
	//var p *byte = &c
	var str_one string
	str_one = "Mi Idea?"
	fmt.Printf("String Program")
	fmt.Println("First is:", str_one)
	fmt.Println("c is:", c)
	//fmt.Printfln("p is:", p)
	//sum := len(str_one)
	//for i := 0; i < sum; i++ {
	for i := range len(str_one) {
		fmt.Printf("char %d is %c \n", i, str_one[i])
	}
}

/* some differences from C Stirng i s a type in go ---in C it is char* so more primite */
