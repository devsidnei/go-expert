package main

import "fmt"

// 5 - Arrays
func main() {

	var languages [3]string
	languages[0] = "Go"
	languages[1] = "Node"
	languages[2] = "PHP"

	for i, v := range languages {
		fmt.Printf("Indice %d => %v \n", i, v)
	}
}
