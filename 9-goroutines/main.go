package main

import (
	"fmt"
	"sync"
)

func saludar(nombre string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hola desde goroutine:", nombre)
}

func main() {
	var wg sync.WaitGroup

	nombres := []string{"Alice", "Bob", "Carlos"}

	for _, nombre := range nombres {
		wg.Add(1)
		go saludar(nombre, &wg)
	}

	wg.Wait()
	fmt.Println("Todas las goroutines terminaron")
}
