package main

import "fmt"

//Exercício 2: Variáveis Globais vs Locais
//Declare uma variável global contador do tipo int e uma variável local
// nome do tipo string dentro da função main(). Exiba ambas no console e
// teste o acesso às variáveis em diferentes escopos.

var (
	contador int = 10
)

func main() {
	var local string = "Teste"

	fmt.Println(contador)
	fmt.Println(local)
}
