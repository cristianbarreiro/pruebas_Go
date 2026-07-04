package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
}

func (p Persona) Saludar() string {
	return fmt.Sprintf("Hola, soy %s y tengo %d años", p.Nombre, p.Edad)
}

func main() {
	p1 := Persona{Nombre: "Ana", Edad: 30}
	p2 := Persona{"Luis", 25}

	fmt.Println(p1.Saludar())
	fmt.Println(p2.Saludar())
	fmt.Println("Nombre:", p1.Nombre, "| Edad:", p1.Edad)
}
