package main

import (
	"fmt"

	"github.com/GustavoCamargodaSilva/api-gateway-golang/matematica" // Atençao a pasta na qual vai ser colocado o go.mod, precisa respeitar a hierarquia de pastas de import.
)

// func com letra maiuscula é exportado
// func com letra minuscula nao é exportado

func main() {
	s := matematica.Soma(10, 20)
	fmt.Println("Resultado: ", s)

	carro := matematica.Carro{Modelo: "Fiat"}
	fmt.Println(carro.Modelo)
}
