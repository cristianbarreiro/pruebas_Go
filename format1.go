package main

import "fmt"

func main() {
	i := 42

	fmt.Println("Here is Println with defaults: i =", i)
	fmt.Printf("Here is Printf with format: i = %d \n", i)
	fmt.Printf("Here is Printf with width 10d format: i = %10d \n", i)
	fmt.Printf("Here is Printf with binary format: i = %b \n", i)
	fmt.Printf("Here is Printf with o octal format: i = %o \n", i)
	fmt.Printf("Here is Printf with x hexadecimal formats: i = %x \n", i)
	fmt.Printf("Here is Printf with O octal formats: i = %O \n", i)

	x := 1000 * i

	fmt.Printf("Here is Printf with g format: i = %g \n", float64(x))
	fmt.Printf("Here is Printf with 10.2g format: i = %10.2g \n", float64(x))
}
