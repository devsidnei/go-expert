package main

import (
	"fmt"

	"github.com/devsidnei/go-expert/7-paking/1/math"
	"github.com/google/uuid"
)

func main() {
	calc := math.NewMath(10, 30)
	fmt.Printf("A Soma é: %d \n", calc.Soma())

	fmt.Println(uuid.New().String())
}
