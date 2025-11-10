package main

import (
	"errors"
	"fmt"
)

func main() {

	//fmt.Println(sum(1, 2))

	//fmt.Println(sum2(51, 2))

	//fmt.Println(sum3(51, 2))

	valor, err := sum3(5, 5) //criando a variavel valor e falando que ela pode retornar um erro ou um valor igual está na chamada do metodo
	if err != nil {          //verificando se há um erro dentro da variavel err
		fmt.Println(err)
	}

	fmt.Println(valor)

}

func sum(a, b int) int { //decaracao do que vai receber e do tipo que a funcao vai retornar
	return a + b
}

func sum2(a, b int) (int, bool) { //Funcao retornando um valor inteiro e um false, é possivel 2 tipos de retornos
	if a+b >= 50 {
		return a + b, true
	}

	return a + b, false
}

func sum3(a, b int) (int, error) { //Funcao retornando um valor inteiro e um false, é possivel 2 tipos de retornos

	if a+b >= 50 {
		return 0, errors.New("A Soma é maior que 50")
	}

	return a + b, nil //nil significa null
}
