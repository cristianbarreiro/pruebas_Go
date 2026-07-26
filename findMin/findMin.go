package main
import (
	"fmt"
	"math/rand"
	"time"
)

func minSlice(d[] float64) float64 {
	var min float64 = 1000.0
	for _, v := range d {
		if v < min {
			min = v
		}
	}
	return min
}

func maxSlice(d[] float64) float64 {
	var max float64 = 0.0
	for _, c := range d {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("size of sice is ")
	size := 0
	fmt.Scanf("%d", &size)
}