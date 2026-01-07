package main

import (
	"encoding/json"
	"os"
)

//Manipular Json

type Conta struct {
	Numero int
	Saldo  int
}

type Conta2 struct {
	Numero int `json:"n"` //-> anotaçao de json para converter corretamente o json para struct EXEMPLO 3
	Saldo  int `json:"s"` // `json:"-"` se deixar o - significa que vou omitir aquela informaçao
}

func main() { //-> json sempre retorna os bytes

	conta := Conta{Numero: 1, Saldo: 200}

	// Convertendo bytes para json

	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}

	println(string(res)) //-> se colocar só o res ele vai dar apenas o endereço de memória

	err = json.NewEncoder(os.Stdout).Encode(conta)

	jsonPuro := []byte(`{"Numero": 2, "Saldo": 200}`) //-> Convertendo Json puro / é chave valor, nomes precisam casar
	var contaX Conta

	if err != nil {
		println(err)
	}
	json.Unmarshal(jsonPuro, &contaX)
	println(contaX.Saldo)

	//------------------------------------------------EXEMPLO 3------------------------------------------------

	jsonPuro2 := []byte(`{"n": 2, "s": 200}`) //-> nome dos json diferente para testar a anotaçao em Conta2
	var contax2 Conta2

	if err != nil {
		println(err)
	}

	json.Unmarshal(jsonPuro2, &contax2)
	println(contax2.Saldo)

}
