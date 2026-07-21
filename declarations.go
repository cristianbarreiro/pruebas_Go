// Package to explore basic types and declarations
package main

import "fmt"

func main() {
	fmt.Println("Will use basic types")
	var myInt int    // simple declaration of int
	anotherInt := 36 // implicit declaration of int
	myFloat := 8.95
	var data1, data2, data3 float32 = 0.5, 1.5, 3.5
	var mData float64
	var unsignedNumber uint

	fmt.Println("Print myInt ", myInt)
	fmt.Println("Print anotherInt ", anotherInt)
	fmt.Println("Print myFloat ", myFloat)
	fmt.Println("Print data ", data1, data2, data3)
	fmt.Println("Print unsignedNumber", unsignedNumber)

	myInt = myInt + anotherInt
	fmt.Println("Print myInt ", myInt)

	unsignedNumber = uint(myInt) // need conversion in GO
	fmt.Println("Print unsignedNumber", unsignedNumber)

	unsignedNumber = uint(-myInt) // a very large number: why???
	fmt.Println("Print unsignedNumber", unsignedNumber)

	mData = float64(data1 + data2 + data3)
	fmt.Println("Print mData", mData)
}
