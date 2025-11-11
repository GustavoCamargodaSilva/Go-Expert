package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c *Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O Cliente %s foi desativado\n", c.Nome)
}

func main() {

	gustavo := Cliente{
		Nome:  "Gustavo",
		Idade: 20,
		Ativo: true,
	}

	//wesley.Ativo = false
	gustavo.Cidade = "Sao Paulo"
	gustavo.Desativar()

	println(gustavo.Ativo)

	//fmt.Printf("O nome é %s, a idade é  %d, ele está ativo? %t", wesley.Nome, wesley.Idade, wesley.Ativo)

}
