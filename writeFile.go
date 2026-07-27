package main

import (
	"fmt"
	"os"
)

func checkFile(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Write File of Int Program locally to intData")
	file, err := os.Create("intData")
	checkFile(err)
	defer file.Close()
	var dataSlice = []int{45, 66, 88, 99, 10}
	for _, v := range dataSlice {
		_, err := fmt.Fprintf(file, "%d\n", v)

		checkFile(err)
	}

	fmt.Println(dataSlice)
}
