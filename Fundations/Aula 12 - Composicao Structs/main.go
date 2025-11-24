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

func main() {

	wesley := Cliente{
		Nome:  "Gustavo",
		Idade: 20,
		Ativo: true,
	}

	wesley.Ativo = false
	wesley.Cidade = "Sao Paulo"

	fmt.Printf("O nome é %s, a idade é  %d, ele está ativo? %t", wesley.Nome, wesley.Idade, wesley.Ativo)

}
