package main

import (
	"io"
	"net/http"
)

func main() {

	req, err := http.Get("https://www.google.com") //Metodo get, post, put e etc
	if err != nil {
		panic(err)
	}

	res, err := io.ReadAll(req.Body) //-> Precisa ler a resposta para poder utilizar
	if err != nil {
		panic(err)
	}
	println(string(res))
	req.Body.Close() //-> importante fechar a requisiçao

}
