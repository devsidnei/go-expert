package main

import "fmt"

func main() {
	var myVar interface{} = "Sidnei Simeão"

	myVar2, ok := myVar.(string)

	fmt.Printf("A interface vazia foi convertida para o tipo %T e a conversão retornou %v \n", myVar2, ok)

	myVar3, ok := myVar.(int)

	if !ok {
		fmt.Printf("Não foi possivel converter a interface vazia para float %v => %v \n", myVar3, ok)

	}

	var myIntVar interface{} = 10

	// Assert que o valor na interface é do tipo int
	intValue, ok := myIntVar.(int)
	if !ok {
		fmt.Println("A asserção de tipo falhou.")
		return
	}

	// Converter o int para float64
	myFloatVar := float64(intValue)

	fmt.Printf("Valor float: %f\n", myFloatVar)
}
