package main

import "fmt"

func main() {
	fmt.Println("Primeira linha")
	defer fmt.Println("Segunda linha") // atraza execução, usado para fechar recursos etc...
	fmt.Println("Terceira linha")
}
