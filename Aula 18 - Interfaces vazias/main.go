package main

import "fmt"

/*
╔══════════════════════════════════════════════════════════════════════════════╗
║                     INTERFACES VAZIAS E PONTEIROS EM GO                      ║
║                                                                              ║
║ Uma INTERFACE VAZIA (interface{}) é uma interface sem nenhum método.         ║
║ Ela pode armazenar qualquer tipo de dado, funcionando como um "tipo curinga" ║
║                                                                              ║
║ RELAÇÃO COM PONTEIROS:                                                      ║
║ - interface{} armazena valores de qualquer tipo                             ║
║ - Pode armazenar ponteiros também: var x interface{} = &minhaVar           ║
║ - Ao recuperar um valor, precisa fazer type assertion                       ║
║                                                                              ║
║ CASOS DE USO:                                                               ║
║ - JSON (marshal/unmarshal): dados podem ter tipos variados                 ║
║ - fmt.Println(...): aceita qualquer tipo                                    ║
║ - Funções genéricas que precisam trabalhar com múltiplos tipos             ║
╚══════════════════════════════════════════════════════════════════════════════╝
*/

func main() {

	// ═══════════════════════════════════════════════════════════════════════════
	// 1. DECLARAÇÃO DE INTERFACES VAZIAS COM VALORES
	// ═══════════════════════════════════════════════════════════════════════════

	// interface{} é um tipo que pode armazenar QUALQUER tipo de dado
	// Não impõe nenhuma restrição sobre qual tipo pode ser atribuído
	// É equivalente a 'Object' em Java ou 'any' em TypeScript

	// Variável x armazena um inteiro dentro de uma interface vazia
	var x interface{} = 10
	// Variável y armazena uma string dentro de uma interface vazia
	var y interface{} = "Hello!!"

	// ═══════════════════════════════════════════════════════════════════════════
	// 2. PASSANDO INTERFACES VAZIAS PARA FUNÇÕES
	// ═══════════════════════════════════════════════════════════════════════════

	// A função showType recebe interface{}, então pode aceitar qualquer tipo
	// Go automaticamente "embrulha" os valores dentro da interface vazia
	showType(x)
	showType(y)

	// ═══════════════════════════════════════════════════════════════════════════
	// 3. EXEMPLO COM PONTEIROS (Bônus educativo)
	// ═══════════════════════════════════════════════════════════════════════════

	// Você também pode armazenar ponteiros em interfaces vazias!
	numero := 42
	var z interface{} = &numero // Armazena um ponteiro para int em interface{}

	showType(z) // Mostra: O tipo da variavel é *int e o valor é 0xc0000...

	// Para acessar o ponteiro, precisamos fazer type assertion
	// Sintaxe: valor.(tipo) retorna (valor, ok)
	if ptr, ok := z.(*int); ok {
		fmt.Printf("Desreferenciando o ponteiro: %d\n", *ptr) // Saída: 42
		*ptr = 100
		fmt.Printf("Modificou através do ponteiro: %d\n", numero) // Saída: 100
	}

	// ═══════════════════════════════════════════════════════════════════════════
	// 4. DIFERENÇA ENTRE ARMAZENAR VALOR vs PONTEIRO
	// ═══════════════════════════════════════════════════════════════════════════

	valor := 50
	var interfaceComValor interface{} = valor     // Copia o valor (50)
	var interfaceComPonteiro interface{} = &valor // Armazena o endereço

	// Modifica a variável original
	valor = 999

	// interfaceComValor continua com 50 (cópia)
	fmt.Printf("Interface com valor: %v\n", interfaceComValor) // 50

	// interfaceComPonteiro reflete a mudança (mesmo endereço)
	if ptr, ok := interfaceComPonteiro.(*int); ok {
		fmt.Printf("Interface com ponteiro: %v\n", *ptr) // 999
	}

}

/*
═══════════════════════════════════════════════════════════════════════════════
FUNÇÃO: showType
═══════════════════════════════════════════════════════════════════════════════

Demonstra o uso de interface{} como parâmetro genérico.
Aceita QUALQUER tipo de valor e exibe seu tipo e seu valor.

PARÂMETRO:
- t interface{} : Pode ser int, string, bool, struct, ponteiro, array, etc.

CODES DE FORMATO:
- %T : Imprime o tipo do valor
- %v : Imprime o valor em formato padrão

IMPORTANTE:
Essa função funciona com qualquer tipo porque:
1. A interface vazia não define nenhum método obrigatório
2. Todos os tipos em Go implementam implicitamente interface{}
3. É o tipo mais genérico disponível em Go
*/
func showType(t interface{}) {
	// %T mostra o tipo dinâmico armazenado em 't'
	// %v mostra o valor armazenado em 't'
	fmt.Printf("O tipo da variavel é %T e o valor é %v\n", t, t)
}

/*
═══════════════════════════════════════════════════════════════════════════════
RELAÇÃO ENTRE INTERFACE{} E PONTEIROS
═══════════════════════════════════════════════════════════════════════════════

QUANDO USAR INTERFACE{} COM PONTEIROS:

1. JSON Deserialization (mais comum):
   - JSON pode conter qualquer tipo de valor
   - interface{} permite receber qualquer coisa do JSON

   var dados interface{}
   json.Unmarshal([]byte(`{"valor": 42}`), &dados)

2. Collections Genéricas:
   - Mapas que podem armazenar valores de diferentes tipos

   var cache map[string]interface{} = map[string]interface{}{
       "idade": 25,
       "nome": "João",
       "usuario": &minhaStruct,
   }

3. Type Assertions para Recuperar:
   - Se você armazenar interface{} = &valor
   - Recupera com: valor, ok := interface{}.(tipo)

   if ptr, ok := dados.(*MinhaStruct); ok {
       // ptr é um *MinhaStruct
   }

COMPARAÇÃO: VALOR vs PONTEIRO em interface{}

Armazenar Valor:
   var x interface{} = 10
   // Go cria uma cópia de 10
   // Tipo: int | Valor: 10

Armazenar Ponteiro:
   numero := 10
   var x interface{} = &numero
   // Go armazena o endereço
   // Tipo: *int | Valor: 0xc0000...
   // Modificações através do ponteiro afetam a variável original

═══════════════════════════════════════════════════════════════════════════════
*/
