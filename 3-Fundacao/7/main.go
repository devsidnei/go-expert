package main

import "fmt"

// 7 - Maps ( hashtable )
func main() {

	languages := map[string]int{}
	languages["GO"] = 1000
	languages["NODE"] = 2000

	fakeLanguages := make(map[string]int)
	fakeLanguages["JAVA"] = -1000
	fakeLanguages["COBOL"] = -2000

	var othersLanguages = map[string]int{
		"PYTHON": 20,
		"PHP":    10,
	}

	// Deletando a entrada "PHP"
	delete(othersLanguages, "PHP")

	// Adicionando a entrada "JAVASCRIPT"
	othersLanguages["JAVASCRIPT"] = 10

	// Imprimindo o mapa atualizado
	fmt.Println(othersLanguages)

	fmt.Println("Outras linguagens", len(othersLanguages))

	// Tambem podem ser percorridos
	for i, v := range othersLanguages {
		fmt.Printf("Linguagem %s => Salário: %d \n", i, v)
	}

	// Usando blank identifier
	for _, v := range othersLanguages {
		fmt.Printf("Salário: %d \n", v)
	}

	// Usando blank identifier
	for i, _ := range othersLanguages {
		fmt.Printf("Linguagem: %s \n", i)
	}
}
