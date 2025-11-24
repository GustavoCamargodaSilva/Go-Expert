package main

import "fmt"

func main() {

	// Declaração de uma variável comum
	// 'a' armazena o valor 10 na memória
	a := 10
	fmt.Println(a) // Saída: 10

	// O operador & (address-of) retorna o ENDEREÇO DE MEMÓRIA de 'a'
	// Mostra onde 'a' está armazenado na memória (exemplo: 0xc000012345)
	fmt.Println(&a) // Saída: 0xc000012345 (endereço em hexadecimal)

	// Declaração explícita de um PONTEIRO
	// *int significa: "ponteiro para um inteiro"
	// ponteiro armazena o ENDEREÇO de 'a', não o valor dele
	var ponteiro *int = &a
	fmt.Println(ponteiro) // Saída: 0xc000012345 (mesmo endereço de &a)

	// Desreferência: usando * na frente de um ponteiro, acessamos o VALOR que ele aponta
	// *ponteiro significa: "vá para o endereço que 'ponteiro' armazena e mude o valor lá"
	// Isso muda o valor de 'a' de 10 para 20
	*ponteiro = 20
	print(a) // Saída: 20 (o valor de 'a' foi alterado através do ponteiro!)

	// Declaração de outro ponteiro usando a sintaxe curta :=
	// b também aponta para o endereço de 'a'
	b := &a
	println(b) // Saída: 0xc000012345 (mesmo endereço, pois aponta para 'a')

	// Desreferência de 'b': obtém o valor que 'b' aponta (que é o valor de 'a')
	println(*b) // Saída: 20 (o valor atual de 'a')

	// Alterando novamente o valor através da desreferência de 'b'
	// Como 'b' aponta para 'a', isso muda o valor de 'a' para 30
	*b = 30
	println(a) // Saída: 30 (o valor de 'a' foi alterado novamente!)

	// RESUMO DE CONCEITOS:
	// - a         : variável que armazena um valor inteiro
	// - &a        : operador & retorna o endereço de memória de 'a'
	// - *int      : tipo de dado que é um ponteiro para inteiro
	// - ponteiro  : variável que armazena um endereço de memória
	// - *ponteiro : desreferência, acessa o valor no endereço armazenado

}
