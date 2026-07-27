package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(d []int) {
	for i := 0 ; i < len(d) - 1 ; i++ {
		for j := 0 ; j < len(d) ; j++ {
			if d[j] > d[j+1] {
				d[j], d[j + 1] = d [j + 1], d[j]
			}
		}
	}
}

func main () {
	fmt.Println(" data of Int random numbers 0 to 10000")
	fmt.Println("time is", time.Now())

	rand.Seed(time Now().UnixNano())
	var size int
	//size = 100000
	fmt.Println("problem size")
	fmt.Scanf("%d", &size)
	dataSlice := make([]int, size)
	for i,_ := range dataSlice {
		dataSlice[i] = rand.Intn(10000)
	}
	//fmt.Println(dataSlice)
	fmt.Println(dataSlice[:20])
	fmt.Println(dataSlice[size - 20:])
	bubbleSort(dataSlice)
	fmt.Println(dataSlice[:20])
	fmt.Println(dataSlice[size - 20:])
	fmt.Println("time is", time.Now())
}