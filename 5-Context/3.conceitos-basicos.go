package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	cep := "01001000" // CEP para consulta

	// Crie um contexto de cancelamento global
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Crie um canal para receber as respostas
	resultChan := make(chan string, 3)

	// Use uma WaitGroup para esperar todas as goroutines terminarem
	var wg sync.WaitGroup

	// Inicie a consulta para cada API em uma goroutine separada
	wg.Add(3)
	go queryAPI(ctx, "ViaCEP", "https://viacep.com.br/ws/"+cep+"/json/", resultChan, &wg)
	go queryAPI(ctx, "Cepaberto", "https://www.cepaberto.com/api/v3/cep?cep="+cep, resultChan, &wg)
	go queryAPI(ctx, "Postmon", "https://api.postmon.com.br/v1/cep/"+cep, resultChan, &wg)

	// Aguarde até que todas as goroutines terminem
	wg.Wait()

	// Feche o canal de resultados
	close(resultChan)

	// Receba a primeira resposta disponível
	response := <-resultChan

	fmt.Printf("Resposta mais rápida: %s\n", response)
}

func queryAPI(ctx context.Context, name, url string, resultChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		// Se o contexto for cancelado antes da conclusão, retorne imediatamente
		return
	default:
		// Faça a consulta à API
		response := simulateAPIRequest(url)
		result := fmt.Sprintf("%s: %s", name, response)
		resultChan <- result
	}
}

func simulateAPIRequest(url string) string {
	// Simula uma solicitação à API com um atraso aleatório
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	return "Resultado da API para " + url
}
