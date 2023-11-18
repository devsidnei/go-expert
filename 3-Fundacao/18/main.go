package main

import "fmt"

func main() {
	// Exempĺo de uso para criar uma lista com elementos de vários tipos
	var lista []interface{}

	// Use a função append para adicionar elementos à slice
	lista = append(lista, "Sidnei Simeão")
	lista = append(lista, 23)
	lista = append(lista, 123.89)
	lista = append(lista, func(a int) int {
		return a
	})

	for i, v := range lista {
		fmt.Printf("O indice %v, é do tipo %T e tem o valor de %v \n", i, v, v)
	}
}
