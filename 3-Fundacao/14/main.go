package main

import "fmt"

// 14 - Interfaces
type Animal interface {
	MakeSound() string
	GetName() string
}

type Dog struct {
	Name string
}

func (d Dog) MakeSound() string {
	return "Woof!"
}

func (d Dog) GetName() string {
	return d.Name
}

type Cat struct {
	Name string
}

func (c Cat) MakeSound() string {
	return "Meow"
}

func (d Cat) GetName() string {
	return d.Name
}

func main() {

	dog := Dog{Name: "Rex"}
	cat := Cat{Name: "whiskers"}

	animals := []Animal{dog, cat}

	for _, animal := range animals {
		fmt.Printf("%s faz o som: %s\n", animal.GetName(), animal.MakeSound())
	}

}
