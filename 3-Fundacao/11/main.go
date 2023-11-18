package main

import "fmt"

type Alien struct {
	Name   string
	Planet string
	Age    int
}

func main() {

	alien := Alien{
		Name:   "Ivan",
		Planet: "Desconhecido",
		Age:    34,
	}

	fmt.Println("Alien:", alien.Name)
	fmt.Println("Planeta:", alien.Planet)
	fmt.Println("Idade:", alien.Age)
}
