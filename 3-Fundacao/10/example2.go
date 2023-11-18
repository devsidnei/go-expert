package main

import "fmt"

func contador() func() int {
	count := 0

	increment := func() int {
		count++
		return count
	}

	return increment
}

func main() {

	c := contador()()

	fmt.Println("Contagem:", c)
	fmt.Println("Contagem:", c)
	fmt.Println("Contagem:", c)
}
