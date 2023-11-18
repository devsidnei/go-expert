package main

import "fmt"

// 9 - Funções variaticas
func main() {

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9))

}

func sum(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}
