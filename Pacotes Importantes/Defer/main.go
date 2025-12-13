package main

import "fmt"

// import (
// 	"io"
// 	"net/http"
// )

// // func main() {

// // 	req, err := http.Get("https://www.google.com") //Metodo get, post, put e etc
// // 	if err != nil {
// // 		panic(err)
// // 	}

// // 	defer req.Body.Close() // -> Defer vai atrasar a execuçao das coisas

// // 	res, err := io.ReadAll(req.Body) //-> Precisa ler a resposta para poder utilizar
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	println(string(res))

// // }

func main() {
	defer fmt.Println("Primeira Linha") //-> Defer vai deixar essa execuçao por unico, sendo os prints como segunda, terceira e primeira.
	fmt.Println("Segunda Linha")
	fmt.Println("Terceira Linha")
}
