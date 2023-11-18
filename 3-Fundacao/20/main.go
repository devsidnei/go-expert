package main

import "math"

type Number interface {
	~int | ~float64
}

func Somar[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomarInteiros(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomarFloats(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return math.Round(soma)
}

func main() {

	salarios := map[string]int{"José": 1000, "Maria": 5000, "Pedro": 2000}

	salarios2 := map[string]float64{"José": 1000.50, "Maria": 5000.50, "Pedro": 2000.50}

	println(SomarFloats(salarios2))
	println(SomarInteiros(salarios))

	println(Somar(salarios))
	println(Somar(salarios2))

}
