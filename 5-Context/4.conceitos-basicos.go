package main

import (
	"encoding/json"
	"fmt"
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

	// Canais para as consultas de CNPJ, ID do cliente e vendedor
	cnpjChan := make(chan string)
	erpIDChan := make(chan string)
	vendorChan := make(chan string)

	// Consulta de CNPJ concorrente
	go func() {
		cnpj := queryCNPJ(order.ClientID)
		cnpjChan <- cnpj
	}()

	// Consulta de ID do cliente no ERP concorrente
	go func() {
		erpID := queryERP(order.ClientID)
		erpIDChan <- erpID
	}()

	// Consulta do vendedor no ERP concorrente
	go func() {
		vendor := queryVendor(order.ClientID)
		vendorChan <- vendor
	}()

	// Aguarde até que todas as consultas sejam concluídas
	cnpj := <-cnpjChan
	erpID := <-erpIDChan
	vendor := <-vendorChan

	// Imprima o pedido apenas se todas as consultas foram bem-sucedidas
	if cnpj != "" && erpID != "" && vendor != "" {
		fmt.Println("Pedido recebido com sucesso!")
		fmt.Printf("ID do Pedido: %d\n", order.ID)
		fmt.Printf("Cliente CNPJ: %s\n", cnpj)
		fmt.Printf("ID do Cliente no ERP: %s\n", erpID)
		fmt.Printf("ID do Vendedor: %s\n", vendor)
	} else {
		fmt.Println("Erro durante as consultas. Pedido não pode ser processado.")
	}
}

// Função fictícia para consulta de CNPJ
func queryCNPJ(clientID string) string {
	// Simule uma consulta demorada ao CNPJ
	time.Sleep(time.Second * 2)
	return "123.456.789/0001-00" // CNPJ fictício
}

// Função fictícia para consulta de ID do cliente no ERP
func queryERP(clientID string) string {
	// Simule uma consulta demorada ao ERP
	time.Sleep(time.Second * 3)
	return "ERP12345" // ID do cliente fictício
}

// Função fictícia para consulta do vendedor no ERP
func queryVendor(clientID string) string {
	// Simule uma consulta demorada ao ERP
	time.Sleep(time.Second * 4)
	return "VEND789" // ID do vendedor fictício
}
