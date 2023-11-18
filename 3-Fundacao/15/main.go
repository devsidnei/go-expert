package main

import "fmt"

var a string = "valor inicial da variavel a"
var b *string = &a

func main() {
	fmt.Printf("Nome da Variável (a): %s\n", "a")
	fmt.Printf("Endereço de Memória (&a): %p\n", &a)
	fmt.Printf("Valor Normal (a): %s\n", a)

	fmt.Println()

	fmt.Printf("Nome da Variável (b): %s\n", "b")
	fmt.Printf("Endereço de Memória (&b): %p\n", &b)
	fmt.Printf("Valor Normal (*b): %s\n", *b)
}
