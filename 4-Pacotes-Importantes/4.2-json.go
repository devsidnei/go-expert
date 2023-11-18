package main

import (
	"encoding/json"
	"fmt"
)

type Phone struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode string `json:"zipcode"`
}

type Customer struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phones    []Phone   `json:"phones"`
	Addresses []Address `json:"addresses"`
}

func main() {
	customer := Customer{
		ID:    1,
		Name:  "João Silva",
		Email: "joao.silva@example.com",
		Phones: []Phone{
			{Type: "celular", Number: "123-456-7890"},
			{Type: "trabalho", Number: "987-654-3210"},
		},
		Addresses: []Address{
			{Street: "123 Main Street", City: "Cidade A", State: "Estado A", Zipcode: "12345"},
		},
	}

	// Marshal a struct Customer para JSON
	jsonData, err := json.MarshalIndent(customer, "", " ")
	if err != nil {
		fmt.Println("Erro ao serializar para JSON:", err)
		return
	}

	// Print o JSON data
	fmt.Println(string(jsonData))
}
