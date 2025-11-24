package main

func main() {
	// -> Loop / Tem apenas o FOR

	for i := 0; i < 10; i++ { //-> laço comum de repetiçao
		println(i)
	}

	numeros := []string{"Um", "Dois", "Tres\n"}

	for n, v := range numeros { // percorrer slices/arrays etc / nesse caso vai mostrar o indice que corresponde aquele valor
		println(n, v)
	}

	for _, n := range numeros { //percorrendo o range ignorando o indice e pegando apenas os valores.
		println(n)
	}

	i := 0 // -> Laço que lembra muito o while
	for i < 10 {
		println(i)
		i++
	}

	for {
		println("Loop, infinito.") //pode ser usado para filas por exemplo
	}
}
