package main

import "fmt"

func main() {

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 8, 7, 6, 5, 4, 3, 2, 1))

}

func sum(numeros ...int) int { //Minha funcao vai entender que posso passar quntos parametros forem desde que ele respeite o tipo int
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
