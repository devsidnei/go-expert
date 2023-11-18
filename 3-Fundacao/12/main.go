package main

import "fmt"

// 12 - Composiçao em Struct
type SpaceShip struct {
	Model string
}

type Alien struct {
	Name      string
	Planet    string
	Age       int
	SpaceShip SpaceShip
}

func main() {
	alien := Alien{
		Name:   "Ivan",
		Planet: "Desconhecido",
		Age:    34,
	}

	alien.SpaceShip.Model = "Xbox"

	fmt.Println("Alien:", alien.Name)
	fmt.Println("Planeta:", alien.Planet)
	fmt.Println("Idade:", alien.Age)
	fmt.Println("Espaçonave:", alien.SpaceShip.Model)
}
