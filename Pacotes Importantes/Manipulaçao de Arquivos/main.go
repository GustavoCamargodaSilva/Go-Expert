package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//-> criaçao do arquivo

	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// -> gravar aquivo

	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo")) //Gravar bytes
	//tamanho, err := f.WriteString("Hello, World!") // escrever hello word no arquivo
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho %d bytes\n", tamanho) //exibir tamanho do arquivo
	f.Close()

	//leitura de arquivo

	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo)) //convertendo para string porque dentro dos arquivos eles sao guardados como bytes.

	//leitura de pouco em pouco abrindo o arquivo

	arquivo2, err := os.Open("arquivo.txt") // -> Abertura de arquivo
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2) // ler o arquivo
	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer) //passando meu reader com arquivo e mandando ler passando meus 10bytes como parametro
		if err != nil {
			break // se houver erro ele vai parar
		}
		fmt.Println(string(buffer[:n])) // converter para string e juntar em um slice, n = posicao
	}

	// como remover um arquivo
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}

}
