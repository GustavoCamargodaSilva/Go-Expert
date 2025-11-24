package main

import "fmt" //importaçao do pacote fmt

type ID int //Com go podemos criar nossos próprios tipos

var (
	a ID      = 1
	b bool    = true
	c int     = 10
	d string  = "Gustavo"
	e float64 = 1.2
)

func main() {

	//Array tem tamanhos fixos.

	var meuArray [3]int //Array com 3 posiçoes inteiros
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30 //atribuiçao de valor para o Array

	fmt.Println(len(meuArray) - 1)         //len está pegando o tamanho do array - 1
	fmt.Println(meuArray[len(meuArray)-1]) //pegando ultima posiçao
	fmt.Println(meuArray[0])               //pegando primeira posiçao que sempre começa com 0

	//PERCORRER ARRAY

	// i = indice
	// v = valor
	// range = percorrer todo o valor

	for i, v := range meuArray {
		fmt.Printf("O valor do indice %d é %d\n", i, v)
	}
}
