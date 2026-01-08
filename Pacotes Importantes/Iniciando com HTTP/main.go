package main

// A palavra "import" permite usar funcionalidades prontas de outros pacotes
// O pacote "net/http" contém ferramentas para criar servidores web
import "net/http"

// A função main() é o ponto de partida do programa
// Quando você executa o programa, o código dentro de main() é executado
func main() {
	// http.HandleFunc() conecta um caminho da URL com uma função que será executada
	// Aqui: quando alguém acessar a raiz "/" do site, a função BuscaCEP será executada
	http.HandleFunc("/", BuscaCEP)

	// http.ListenAndServe() inicia o servidor web
	// ":8080" = servidor escuta na porta 8080 (você acessará digitando localhost:8080 no navegador)
	// nil = significa que usaremos as configurações padrão do servidor
	http.ListenAndServe(":8080", nil)
}

// BuscaCEP é uma função que responde às requisições HTTP
// Toda requisição HTTP necessita de dois parâmetros:
// - w (http.ResponseWriter): objeto para enviar a resposta ao navegador
// - r (*http.Request): objeto que contém informações sobre o pedido recebido
func BuscaCEP(w http.ResponseWriter, r *http.Request) {
	// w.Write() envia dados de resposta para o cliente (navegador)
	// []byte("Hello, World!") converte o texto em um formato que o navegador entende
	// O cliente verá "Hello, World!" exibido na página
	w.Write([]byte("Hello, World!"))
}
