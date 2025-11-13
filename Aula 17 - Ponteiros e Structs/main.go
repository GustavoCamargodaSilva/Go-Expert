package main

import "fmt"

// Definição da struct Cliente
// Uma struct é um tipo de dado que agrupa múltiplos campos relacionados
type Cliente struct {
	nome string
}

// Definição da struct Conta
// Esta struct será usada com ponteiros para demonstrar modificações
type Conta struct {
	saldo int
}

// MÉTODO COM PONTEIRO: (c *Conta)
// Quando o receptor é um PONTEIRO, o método pode MODIFICAR a struct original
//
// Explicação:
// - *Conta significa que o método recebe um ponteiro para Conta
// - Dentro do método, c.saldo += valor modifica o SALDO ORIGINAL
// - Essa é a maneira idiomática em Go para métodos que modificam dados
//
// Nota: Go permite usar c.saldo mesmo que c seja um ponteiro (syntactic sugar)
func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

// MÉTODO SEM PONTEIRO: (c Cliente)
// Quando o receptor é um VALOR, o método recebe uma CÓPIA da struct
// Qualquer modificação feita aqui NÃO afeta a struct original
//
// Explicação:
// - c Cliente significa que o método recebe uma CÓPIA de Cliente
// - c.nome = "Gustavo" modifica apenas a CÓPIA local
// - O valor original de cliente.nome permanece inalterado
//
// Isso é útil quando você quer fazer operações sem efeitos colaterais
func (c Cliente) andou() {
	c.nome = "Gustavo"
	fmt.Printf("O Cliente %v andou \n", c.nome)
}

// FACTORY FUNCTION: Retorna um ponteiro para uma nova Conta
//
// Por que retornar um ponteiro?
// - Quando você quer que o chamador possa modificar a struct retornada
// - Struct Conta será usada com método que tem receptor ponteiro (simular)
// - Padrão comum em Go: factory functions retornam ponteiros
//
// Nota: &Conta{saldo: 0} cria uma nova Conta e retorna seu endereço
func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func main() {

	// Criando uma instância de Cliente
	cliente := Cliente{
		nome: "Wesley",
	}

	// Chamando método sem ponteiro
	// O método recebe uma CÓPIA, então modifica apenas a cópia local
	cliente.andou()

	// O valor original de cliente.nome continua "Wesley" (não foi alterado)
	// A modificação para "Gustavo" aconteceu apenas dentro do método andou()
	fmt.Printf("O valor da Struct com nome %v \n", cliente.nome) // Saída: Wesley

	// Criando uma instância de Conta
	conta := Conta{
		saldo: 100,
	}

	// Chamando método com ponteiro
	// Mesmo sem usar & aqui, Go automaticamente passa o ponteiro
	// porque o método simular foi definido com receptor *Conta
	// O método MODIFICA o saldo original
	conta.simular(200) // Saída: 300

	// O saldo foi modificado para 300 pela chamada acima
	fmt.Printf("%v", conta.saldo) // Saída: 300

	// RESUMO DO QUE APRENDEMOS:
	//
	// 1. Método com valor (c Cliente):
	//    - Recebe uma CÓPIA da struct
	//    - Modificações NÃO afetam o original
	//    - cliente.andou() não alterou cliente.nome
	//
	// 2. Método com ponteiro (c *Conta):
	//    - Recebe um PONTEIRO para a struct
	//    - Modificações AFETAM o original
	//    - conta.simular(200) alterou conta.saldo
	//
	// 3. Go é "inteligente" com sintaxe:
	//    - Você chama conta.simular() (não &conta.simular())
	//    - Go automaticamente passa o ponteiro quando necessário
	//    - Mas conceitualmente, é um ponteiro mesmo assim
	//
	// 4. Regra prática:
	//    - Use ponteiro como receptor quando: quer modificar a struct
	//    - Use valor como receptor quando: quer apenas ler dados

}
