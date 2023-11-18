package veiculo

import "strconv"

type Carro struct {
	modelo string
	marca  string
	ano    int
}

func (c *Carro) SetModelo(modelo string) {
	c.modelo = modelo
}

func (c *Carro) SetMarca(marca string) {
	c.marca = marca
}

func (c *Carro) SetAno(ano int) {
	c.ano = ano
}

func ImprimirDados(c *Carro) string {
	return "Modelo:" + c.modelo + "\n Marca:" + c.marca + "\n Ano:" + strconv.Itoa(c.ano) + "\n"
}
