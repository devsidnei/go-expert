package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type Order struct {
	ID       int    `json:"id"`
	ClientID string `json:"client_id"`
}

type ClientInfo struct {
	CNPJ string `json:"cnpj"`
}

type ERPInfo struct {
	ClientID string `json:"client_id"`
	VendorID string `json:"vendor_id"`
}

func main() {
	// Pedido fictício recebido
	orderJSON := `{"id": 1, "client_id": "123456789"}`
	var order Order
	if err := json.Unmarshal([]byte(orderJSON), &order); err != nil {
		fmt.Println("Erro ao decodificar o pedido:", err)
		return
	}

	// Canal para receber os resultados de todas as goroutines
	resultChan := make(chan string, 3)

	// Use uma WaitGroup para esperar até que todas as goroutines sejam concluídas
	var wg sync.WaitGroup

	// Inicie as consultas concorrentes
	wg.Add(3)
	go queryCNPJ(order.ClientID, resultChan, &wg)
	go queryERP(order.ClientID, resultChan, &wg)
	go queryVendor(order.ClientID, resultChan, &wg)

	// Espere até que todas as goroutines terminem
	wg.Wait()

	// Feche o canal de resultados
	close(resultChan)

	// Recolha os resultados das goroutines
	var results []string
	for result := range resultChan {
		results = append(results, result)
	}

	// Verifique se todas as consultas foram bem-sucedidas e imprima o pedido
	if len(results) == 3 {
		fmt.Println("Pedido recebido com sucesso!")
		fmt.Printf("ID do Pedido: %d\n", order.ID)
		fmt.Printf("Cliente CNPJ: %s\n", results[0])
		fmt.Printf("ID do Cliente no ERP: %s\n", results[1])
		fmt.Printf("ID do Vendedor: %s\n", results[2])
	} else {
		fmt.Println("Erro durante as consultas. Pedido não pode ser processado.")
	}
}

// Função fictícia para consulta de CNPJ
func queryCNPJ(clientID string, resultChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simule uma consulta demorada ao CNPJ
	time.Sleep(time.Second * 2)
	cnpj := "123.456.789/0001-00" // CNPJ fictício
	resultChan <- cnpj
}

// Função fictícia para consulta de ID do cliente no ERP
func queryERP(clientID string, resultChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simule uma consulta demorada ao ERP
	time.Sleep(time.Second * 3)
	erpID := "ERP12345" // ID do cliente fictício
	resultChan <- erpID
}

// Função fictícia para consulta do vendedor no ERP
func queryVendor(clientID string, resultChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simule uma consulta demorada ao ERP
	time.Sleep(time.Second * 4)
	vendor := "VEND789" // ID do vendedor fictício
	resultChan <- vendor
}
