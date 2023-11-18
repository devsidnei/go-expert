package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*f, err := os.Create("arquivo.txt")

	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("Hello World"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes \n", tamanho)

	f.Close()
	*/

	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)

	for {
		tamanho, err := reader.Read(buffer)
		if err != nil {
			fmt.Println("Erro ao ler arquivo:", err)
			break
		}
		fmt.Println("Bytes lidos:", tamanho, "Conteúdo:", string(buffer[:tamanho]))
	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
