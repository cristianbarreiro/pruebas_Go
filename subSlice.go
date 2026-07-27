package main

import (
	"fmt"
)

func main() {

	data := []int{3, 4, 5, 6, 7} // declare and initialize slice

	data2 := data[:2]
	data3 := data[0:2]
	data4 := data[2:4]
	data5 := data[2:]

	fmt.Println("Demo sub Slicing")

	fmt.Println("Original data:")
	for _, v := range data {
		fmt.Println(v)
	}

	fmt.Println("data[:2]")
	for _, v := range data2 {
		fmt.Println(v)
	}

	fmt.Println("data[0:2]")
	for _, v := range data3 {
		fmt.Println(v)
	}

	fmt.Println("data[2:4]")
	for _, v := range data4 {
		fmt.Println(v)
	}

	fmt.Println("data[2:]")
	for _, v := range data5 {
		fmt.Println(v)
	}
}
