package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	wesley := Cliente{
		Nome:  "Gustavo",
		Idade: 20,
		Ativo: true,
	}

	wesley.Ativo = false

	fmt.Printf("O nome é %s, a idade é  %d, ele está ativo? %t", wesley.Nome, wesley.Idade, wesley.Ativo)

}
