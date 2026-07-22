package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Printf("%d \t", rand.Intn(6)+1)
	}
	fmt.Println("BYE")
}
