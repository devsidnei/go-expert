package main

import (
	"fmt"
)

var users = map[int]string{
	1: "Sidnei",
	2: "Fábio",
	3: "Ivan",
	4: "Lontra",
}

func main() {
	for i := range users {
		if isAlien, name := userIsAlien(i); isAlien {
			fmt.Printf("O Usuário %s é um Alien \n", name)
		} else {
			fmt.Printf("O Usuário %s não é um Alien \n", name)
		}

	}
}

func userIsAlien(id int) (bool, string) {
	alienName, isAlien := users[id]
	return isAlien && alienName == "Ivan", alienName
}
