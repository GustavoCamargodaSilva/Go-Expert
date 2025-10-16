package main

import "fmt" //importaçao do pacote fmt

type ID int //Com go podemos criar nossos próprios tipos

var (
	a ID = 1
)

func main() {

	println(a)
	fmt.Printf("O TIPO DE A É %T", a)

}
