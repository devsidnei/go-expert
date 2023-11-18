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

	jsonPuro := []byte(`	
		  {
			"id": 1,
			"name": "João Silva",
			"email": "joao.silva@example.com",
			"phones": [
			  {
				"type": "celular",
				"number": "123-456-7890"
			  },
			  {
				"type": "trabalho",
				"number": "987-654-3210"
			  }
			],
			"addresses": [
			  {
				"street": "123 Main Street",
				"city": "Cidade A",
				"state": "Estado A",
				"zipcode": "12345"
			  }
			]
		  }`)

	var customer Customer

	err := json.Unmarshal(jsonPuro, &customer)
	if err != nil {
		fmt.Println("Erro ao fazer a análise JSON:", err)
		return
	}

	fmt.Printf("ID: %d\n", customer.ID)
	fmt.Printf("Name: %s\n", customer.Name)
	fmt.Printf("Email: %s\n", customer.Email)

	fmt.Println("Phones:")
	for _, phone := range customer.Phones {
		fmt.Printf("  Type: %s, Number: %s\n", phone.Type, phone.Number)
	}

	fmt.Println("Addresses:")
	for _, address := range customer.Addresses {
		fmt.Printf("  Street: %s, City: %s, State: %s, Zipcode: %s\n", address.Street, address.City, address.State, address.Zipcode)
	}

}
