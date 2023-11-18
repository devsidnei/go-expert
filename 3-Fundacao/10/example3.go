package main

import "fmt"

func contador(fn func(int, int) int, a, b int) func() int {
	result := 0
	return func() int {
		result += fn(a, b)
		return result
	}
}

func main() {

	c := contador(func(a, b int) int { return a + b }, 41, 9)

	fmt.Println("Contagem:", c())
	fmt.Println("Contagem:", c())
	fmt.Println("Contagem:", c())
}
