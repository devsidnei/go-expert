package main

import (
	"context"
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
	fmt.Println("Programa iniciado.")

	// Pedido fictício recebido
	orderJSON := `{"id": 1, "client_id": "123456789"}`

	var order Order
	if err := json.Unmarshal([]byte(orderJSON), &order); err != nil {
		fmt.Println("Erro ao decodificar o pedido:", err)
		return
	}

	// Crie um contexto com timeout de 3 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Canal para receber os resultados de todas as goroutines
	resultChan := make(chan string, 3)

	// Use uma WaitGroup para esperar até que todas as goroutines sejam concluídas
	var wg sync.WaitGroup

	// Inicie as consultas concorrentes
	fmt.Println("Iniciando consultas...")
	wg.Add(3)
	go queryCNPJ(ctx, order.ClientID, resultChan, &wg)
	go queryERP(ctx, order.ClientID, resultChan, &wg)
	go queryVendor(ctx, order.ClientID, resultChan, &wg)

	// Espere até que todas as goroutines terminem ou o contexto expire
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Todas as consultas foram concluídas.")
		// Todas as goroutines foram concluídas com sucesso
		close(resultChan)
	case <-ctx.Done():
		// O contexto atingiu o timeout de 3 segundos
		fmt.Println("Timeout: As consultas não foram concluídas a tempo.")
		return
	}

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
func queryCNPJ(ctx context.Context, clientID string, resultChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	startTime := time.Now()
	fmt.Printf("Iniciando consulta de CNPJ (início em %s)...\n", startTime.Format("2006-01-02 15:04:05.000000"))
	select {
	case <-ctx.Done():
		fmt.Println("Consulta de CNPJ cancelada.")
		return // A goroutine é encerrada quando o contexto é cancelado
	default:
		// Simule uma consulta demorada ao CNPJ
		time.Sleep(time.Second * 2)
		cnpj := "123.456.789/0001-00" // CNPJ fictício
		resultChan <- cnpj
		endTime := time.Now()
		fmt.Printf("Consulta de CNPJ concluída (tempo de execução: %s).\n", endTime.Sub(startTime))
	}
}

// Função fictícia para consulta de ID do cliente no ERP
func queryERP(ctx context.Context, clientID string, resultChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	startTime := time.Now()
	fmt.Printf("Iniciando consulta de ID do cliente no ERP (início em %s)...\n", startTime.Format("2006-01-02 15:04:05.000000"))
	select {
	case <-ctx.Done():
		fmt.Println("Consulta de ID do cliente no ERP cancelada.")
		return // A goroutine é encerrada quando o contexto é cancelado
	default:
		// Simule uma consulta demorada ao ERP
		time.Sleep(time.Second * 3)
		erpID := "ERP12345" // ID do cliente fictício
		resultChan <- erpID
		endTime := time.Now()
		fmt.Printf("Consulta de ID do cliente no ERP concluída (tempo de execução: %s).\n", endTime.Sub(startTime))
	}
}

// Função fictícia para consulta do vendedor no ERP
func queryVendor(ctx context.Context, clientID string, resultChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	startTime := time.Now()
	fmt.Printf("Iniciando consulta do vendedor no ERP (início em %s)...\n", startTime.Format("2006-01-02 15:04:05.000000"))
	select {
	case <-ctx.Done():
		fmt.Println("Consulta do vendedor no ERP cancelada.")
		return // A goroutine é encerrada quando o contexto é cancelado
	default:
		// Simule uma consulta demorada ao ERP
		time.Sleep(time.Second * 4)
		vendor := "VEND789" // ID do vendedor fictício
		resultChan <- vendor
		endTime := time.Now()
		fmt.Printf("Consulta do vendedor no ERP concluída (tempo de execução: %s).\n", endTime.Sub(startTime))
	}
}
