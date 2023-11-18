package main

import "fmt"

// 6 - Slices
func main() {

	/*
		var languages []string
		languages[0] = "GO"
		languages[1] = "NODE"
		languages[2] = "PHP"
		languages[3] = "RUBY"
		languages[4] = "C"
		languages[5] = "C++"
		languages[6] = "PYTHON"
		languages[7] = "SCALA"
	*/

	// Short Sintax
	languages := []string{"GO", "NODE", "PHP", "RUBY", "C", "C++", "PYTHON", "SCALA"}

	// tamanho, capacidade, valores conforme declarados
	fmt.Printf("len=%d, cap=%d, val=%v \n", len(languages), cap(languages), languages)

	// tamanho, capacidade, valores [:2] somente 2 primeiros indices ( não altera a capacidade )
	fmt.Printf("len=%d, cap=%d, val=%v \n", len(languages[:2]), cap(languages[:2]), languages[:2])

	// tamanho, capacidade, valores [2:] apos 2 primeiros indices ( altera a capacidade ( não se preocupe, se precisar ele dobra capacidade automaticamente ))
	fmt.Printf("len=%d, cap=%d, val=%v \n", len(languages[2:]), cap(languages[2:]), languages[2:])

	// inclui um item mas não é linguagem de programação só usei para causar intrigas
	languages = append(languages, "HTML")

	// ao ultrapassar a quantidade inicial, ele dobra de tamanho
	fmt.Printf("len=%d, cap=%d, val=%v \n", len(languages), cap(languages), languages)

	/*
		len=8, cap=8, val=[GO NODE PHP RUBY C C++ PYTHON SCALA]
		len=2, cap=8, val=[GO NODE]
		len=6, cap=6, val=[PHP RUBY C C++ PYTHON SCALA]
		len=9, cap=16, val=[GO NODE PHP RUBY C C++ PYTHON SCALA HTML]
	*/

}
