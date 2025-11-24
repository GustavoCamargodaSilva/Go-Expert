package main

import "fmt"

func main() {

	salarios := map[string]int{"Gustavo": 1000, "Joao": 2000, "Maria": 3000} //Maps de chave e valor como nas outras linguagens
	fmt.Println(salarios)
	fmt.Println(salarios["Gustavo"])
	delete(salarios, "Maria")
	fmt.Println(salarios)
	salarios["Gus"] = 5000
	fmt.Println(salarios)

	// sal := make(map[string]int)
	// sal1 := map[string]int{}

	//PERCORRENDO O MAP

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}

}
