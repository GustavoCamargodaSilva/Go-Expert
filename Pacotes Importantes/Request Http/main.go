// Pacote main: todo programa Go executável precisa ter esse pacote
package main

// import: importa bibliotecas externas que serão usadas no programa
import (
	"encoding/json" // Biblioteca para converter dados em formato JSON
	"fmt"           // Biblioteca para exibir mensagens formatadas
	"io"            // Biblioteca para leitura de dados (entrada/saída)
	"net/http"      // Biblioteca para fazer requisições HTTP (comunicar com servidores web)
	"os"            // Biblioteca para trabalhar com o sistema operacional (argumentos, criar arquivos, etc)
)

// type ViaCEP: define uma estrutura (type) que representa os dados retornados pela API ViaCEP
// Uma estrutura é como um molde que organiza informações relacionadas
// Os nomes em “ são tags que indicam como os dados JSON devem ser mapeados para cada campo
type ViaCEP struct {
	Cep         string `json:"cep"`         // Código postal (CEP)
	Logradouro  string `json:"logradouro"`  // Rua ou avenida
	Complemento string `json:"complemento"` // Informações adicionais do endereço
	Unidade     string `json:"unidade"`     // Unidade ou lote
	Bairro      string `json:"bairro"`      // Nome do bairro
	Localidade  string `json:"localidade"`  // Cidade
	Uf          string `json:"uf"`          // Estado (sigla)
	Estado      string `json:"estado"`      // Nome do estado
	Regiao      string `json:"regiao"`      // Região do Brasil
	Ibge        string `json:"ibge"`        // Código IBGE da cidade
	Gia         string `json:"gia"`         // Código GIA
	Ddd         string `json:"ddd"`         // Código de área telefônico
	Siafi       string `json:"siafi"`       // Código SIAFI
}

// func main(): função principal que é executada automaticamente quando o programa inicia
func main() {

	// for...range: faz um loop (repetição) através de cada URL passada como argumento de linha de comando
	// os.Args[1:] significa: pegue todos os argumentos a partir do segundo (índice 1) até o final
	// O primeiro argumento (índice 0) é sempre o nome do programa executável
	for _, url := range os.Args[1:] {
		// Exibe a URL atual no console (para visualizar qual URL está sendo processada)
		println(url)

		// http.Get(): faz uma requisição GET para a URL especificada
		// GET é um tipo de requisição HTTP para recuperar informações de um servidor
		// Retorna uma resposta (req) ou um erro (err)
		req, err := http.Get(url) // poderia ter colocado o link da viaCEP e concatenar variáveis e etc

		// Verifica se houve erro na requisição
		if err != nil {
			// Se houver erro, exibe a mensagem de erro na tela e continua para a próxima URL
			fmt.Fprintf(os.Stderr, "Erro na requisiçao: %v\n", err)
		}

		// defer: agenda a ação para ser executada ao final da função
		// Close(): fecha a conexão com o servidor após terminar de usar
		// Isso libera recursos do computador para serem usados por outros programas
		defer req.Body.Close()

		// io.ReadAll(): lê TODOS os dados retornados pelo servidor
		// req.Body contém o conteúdo da resposta
		// res será uma fatia de bytes (dados brutos) com a resposta
		res, err := io.ReadAll(req.Body)

		// Verifica se houve erro ao ler a resposta
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler a resposta: %v\n", err)
		}

		// var data ViaCEP: cria uma variável chamada 'data' do tipo ViaCEP
		// Isso prepara um "espaço" para guardar os dados que vêm da API
		var data ViaCEP

		// json.Unmarshal(): converte dados em formato JSON (texto estruturado) para a estrutura Go
		// res: os dados brutos recebidos do servidor (em JSON)
		// &data: o endereço da variável onde os dados devem ser armazenados (& significa "endereço de")
		err = json.Unmarshal(res, &data)

		// Verifica se houve erro na conversão JSON
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao converter: %v\n", err)
		}

		// Exibe os dados convertidos no console de forma formatada
		fmt.Println(data)

		// os.Create(): cria um novo arquivo chamado "cidade.txt" no diretório atual
		// Se o arquivo já existir, ele será sobrescrito (apagado e recriado)
		// 'file' armazena a referência ao arquivo aberto
		file, err := os.Create("cidade.txt")

		// Verifica se houve erro ao criar o arquivo
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao gravar arquivo: %v\n", err)
		}

		// defer: agenda o fechamento do arquivo para quando a função terminar
		// Close(): fecha o arquivo para que ele possa ser usado por outros programas
		defer file.Close()

		// file.WriteString(): escreve texto no arquivo
		// data.Localidade: escreve apenas o nome da cidade no arquivo
		_, err = file.WriteString(data.Localidade)

	}

}
