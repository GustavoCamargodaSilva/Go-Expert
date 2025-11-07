package main

import "fmt"

func main() {

	s := []int{2, 4, 6, 8, 10}                           //declaraçao do slice int e ja atribuindo valores
	fmt.Printf("len= %d cap=%d %v\n", len(s), cap(s), s) //Aqui estou imprimindo len = tamanho do slice - cap = capacidade do slice e s que sao os valores de dentro

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0]) //estou falando para ele imprimir 0 posicoes

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4]) //o :4 ele vai mostrar ATÉ 4 posiçoes

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:]) // o :2 quer dizer que ele vai começar a partir da 2 posiçao - Capacidade diminuida porque ignoramos o começo

	s = append(s, 110) //adicionando o 110 na ponta

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
