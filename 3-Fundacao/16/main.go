package main

import "fmt"

func imcrement(number *int) int {
	*number++
	return *number
}

func decrement(number int) int {
	number--
	return number
}

// Ponteiros
func main() {
	a := 1

	imcrement(&a)
	decrement(a)
	imcrement(&a)
	decrement(a)
	imcrement(&a)
	decrement(a)

	fmt.Println(a)
}
