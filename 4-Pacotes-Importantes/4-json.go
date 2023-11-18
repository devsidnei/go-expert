package main

import (
	"encoding/json"
	"os"
)

// pode-se usar tags ( especie de anotações key=value )
type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {

	// Converter Struc em Json
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)

	if err != nil {
		println(err)
	}

	println(string(res))

	// Retornar um json ( não armazena em variave, já despeja o conteudo no terminal )
	err = json.NewEncoder(os.Stdout).Encode(conta)

	if err != nil {
		println(err)
	}

	// Json from string, sempre representado por um mapa de bytes, converter em Struct
	jsonPuro := []byte(`{"Numero": 2, "Saldo": 200}`)
	var conta2 Conta

	err = json.Unmarshal(jsonPuro, &conta2)
	if err != nil {
		println(err)
	}

	println(string(jsonPuro))

	jsonPuro2 := []byte(`{"n":2, "s":200}`)

	err = json.Unmarshal(jsonPuro2, &conta2)
	if err != nil {
		println(err)
	}

	println(conta2.Saldo)

}
