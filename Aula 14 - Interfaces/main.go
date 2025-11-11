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

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

func (c *Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O Cliente %s foi desativado\n", c.Nome)
}

type Pessoa interface {
	//interface em go permite apenas implementar METODOS
	Desativar()
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {

	gustavo := Cliente{
		Nome:  "Gustavo",
		Idade: 20,
		Ativo: true,
	}

	fmt.Println(gustavo)
	gustavo.Desativar()
	fmt.Println(gustavo)

	minhaEmpresa := Empresa{
		Nome: "Google",
	}

	Desativacao(minhaEmpresa)

}
