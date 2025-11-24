package main

import "fmt"

/*
╔══════════════════════════════════════════════════════════════════════════════╗
║                        CONDICIONAIS EM GOLANG                               ║
║                                                                              ║
║ Estruturas de controle de fluxo que permitem executar código diferente      ║
║ baseado em condições booleanas (true ou false).                             ║
║                                                                              ║
║ TIPOS DISPONÍVEIS:                                                          ║
║ 1. if / else if / else                                                      ║
║ 2. if com inicialização                                                     ║
║ 3. switch / case / default                                                  ║
║ 4. switch sem argumento                                                     ║
║ 5. Operador ternário (não existe em Go, use if/else)                       ║
╚══════════════════════════════════════════════════════════════════════════════╝
*/

func main() {

	// ═══════════════════════════════════════════════════════════════════════════
	// 1. IF / ELSE IF / ELSE - Simples condicional
	// ═══════════════════════════════════════════════════════════════════════════

	fmt.Println("═══ 1. IF / ELSE IF / ELSE ═══")

	a := 1
	b := 2

	// Estrutura básica: if condição { } else { }
	if a > b {
		fmt.Println("a é maior que b")
	} else if a == b {
		fmt.Println("a é igual a b")
	} else {
		fmt.Println("a é menor que b") // Essa linha executa
	}

	// ═══════════════════════════════════════════════════════════════════════════
	// 2. IF COM INICIALIZAÇÃO - Declarar variável antes da condição
	// ═══════════════════════════════════════════════════════════════════════════

	fmt.Println("\n═══ 2. IF COM INICIALIZAÇÃO ═══")

	// Sintaxe: if variavel := valor; condição { }
	// A variável é LOCAL apenas no escopo do if/else
	// Muito útil para erros e valores calculados

	if resultado := a + b; resultado > 5 {
		fmt.Println("Resultado maior que 5:", resultado)
	} else {
		fmt.Println("Resultado menor ou igual a 5:", resultado)
	}

	// resultado não existe aqui fora do if
	// fmt.Println(resultado) // ERRO: undefined

	// ═══════════════════════════════════════════════════════════════════════════
	// 3. SWITCH / CASE / DEFAULT - Múltiplas opções
	// ═══════════════════════════════════════════════════════════════════════════

	fmt.Println("\n═══ 3. SWITCH / CASE / DEFAULT ═══")

	// Sintaxe: switch variavel { case valor1: case valor2: default: }
	// Compara variável contra múltiplos valores
	// Em Go, não precisa de 'break' (sai automaticamente após case)

	dia := 3

	switch dia {
	case 1:
		fmt.Println("Segunda-feira")
	case 2:
		fmt.Println("Terça-feira")
	case 3:
		fmt.Println("Quarta-feira") // Executa esse
	case 4:
		fmt.Println("Quinta-feira")
	case 5:
		fmt.Println("Sexta-feira")
	default:
		fmt.Println("Fim de semana")
	}

	// ═══════════════════════════════════════════════════════════════════════════
	// 4. SWITCH COM MÚLTIPLOS VALORES POR CASE
	// ═══════════════════════════════════════════════════════════════════════════

	fmt.Println("\n═══ 4. SWITCH COM MÚLTIPLOS VALORES ═══")

	// Sintaxe: case valor1, valor2, valor3:
	// Executa se a variável combinar com QUALQUER um dos valores

	mes := 2

	switch mes {
	case 12, 1, 2:
		fmt.Println("Verão (dezembro, janeiro, fevereiro)")
	case 3, 4, 5:
		fmt.Println("Outono (março, abril, maio)")
	case 6, 7, 8:
		fmt.Println("Inverno (junho, julho, agosto)")
	case 9, 10, 11:
		fmt.Println("Primavera (setembro, outubro, novembro)")
	default:
		fmt.Println("Mês inválido")
	}

	// ═══════════════════════════════════════════════════════════════════════════
	// 5. SWITCH SEM ARGUMENTO - Equivalente a if/else if/else
	// ═══════════════════════════════════════════════════════════════════════════

	fmt.Println("\n═══ 5. SWITCH SEM ARGUMENTO ═══")

	// Sintaxe: switch { case condição: }
	// Útil quando cada case tem uma condição diferente
	// Mais legível que múltiplos if/else if/else

	idade := 25

	switch {
	case idade < 13:
		fmt.Println("Criança")
	case idade < 18:
		fmt.Println("Adolescente")
	case idade < 65:
		fmt.Println("Adulto") // Executa esse
	default:
		fmt.Println("Idoso")
	}

	// ═══════════════════════════════════════════════════════════════════════════
	// 6. SWITCH COM INICIALIZAÇÃO
	// ═══════════════════════════════════════════════════════════════════════════

	fmt.Println("\n═══ 6. SWITCH COM INICIALIZAÇÃO ═══")

	// Sintaxe: switch variavel := valor; variavel { case: }
	// Declara variável LOCAL apenas no escopo do switch

	switch tipo := "admin"; tipo {
	case "admin":
		fmt.Println("Acesso total") // Executa esse
	case "user":
		fmt.Println("Acesso limitado")
	case "guest":
		fmt.Println("Sem acesso")
	default:
		fmt.Println("Tipo desconhecido")
	}

	// ═══════════════════════════════════════════════════════════════════════════
	// 7. SWITCH COM TIPOS (Type Switch) - Útil com interfaces{}
	// ═══════════════════════════════════════════════════════════════════════════

	fmt.Println("\n═══ 7. SWITCH COM TIPOS (Type Switch) ═══")

	// Sintaxe: switch valor := variável.(type) { case tipo1: case tipo2: }
	// Verifica o TIPO da variável, não o valor

	var valor interface{} = 42

	switch v := valor.(type) {
	case int:
		fmt.Printf("É um inteiro: %d\n", v) // Executa esse
	case string:
		fmt.Printf("É uma string: %s\n", v)
	case bool:
		fmt.Printf("É um booleano: %v\n", v)
	default:
		fmt.Printf("Tipo desconhecido: %T\n", v)
	}

	// ═══════════════════════════════════════════════════════════════════════════
	// 8. FALLTHROUGH - Continuar para o próximo case
	// ═══════════════════════════════════════════════════════════════════════════

	fmt.Println("\n═══ 8. FALLTHROUGH ═══")

	// Sintaxe: case valor: ... fallthrough
	// Continua executando o próximo case (raro usar)
	// Diferente de linguagens como C/Java que usam break

	numero := 2

	switch numero {
	case 1:
		fmt.Println("Um")
	case 2:
		fmt.Println("Dois")
		fallthrough // Continua para o próximo case
	case 3:
		fmt.Println("Três") // Também executa
	default:
		fmt.Println("Outro")
	}

	// ═══════════════════════════════════════════════════════════════════════════
	// 9. OPERADORES LÓGICOS - AND (&&), OR (||), NOT (!)
	// ═══════════════════════════════════════════════════════════════════════════

	fmt.Println("\n═══ 9. OPERADORES LÓGICOS ═══")

	x := 15
	y := 10

	// AND (&&) - Ambas condições devem ser verdadeiras
	if x > 10 && y < 20 {
		fmt.Println("AND: Ambas verdadeiras") // Executa
	}

	// OR (||) - Pelo menos uma condição deve ser verdadeira
	if x > 20 || y < 15 {
		fmt.Println("OR: Pelo menos uma é verdadeira") // Executa
	}

	// NOT (!) - Nega a condição
	if !(x < 10) {
		fmt.Println("NOT: x não é menor que 10") // Executa
	}

	// ═══════════════════════════════════════════════════════════════════════════
	// 10. OPERADORES COMPARAÇÃO - ==, !=, <, >, <=, >=
	// ═══════════════════════════════════════════════════════════════════════════

	fmt.Println("\n═══ 10. OPERADORES COMPARAÇÃO ═══")

	p := 5
	q := 5

	fmt.Printf("p == q: %v (igualdade)\n", p == q)      // true
	fmt.Printf("p != q: %v (desigualdade)\n", p != q)   // false
	fmt.Printf("p < q: %v (menor que)\n", p < q)        // false
	fmt.Printf("p > q: %v (maior que)\n", p > q)        // false
	fmt.Printf("p <= q: %v (menor ou igual)\n", p <= q) // true
	fmt.Printf("p >= q: %v (maior ou igual)\n", p >= q) // true

	// ═══════════════════════════════════════════════════════════════════════════
	// 11. VERIFICAÇÃO DE NIL - Comum com ponteiros e interfaces
	// ═══════════════════════════════════════════════════════════════════════════

	fmt.Println("\n═══ 11. VERIFICAÇÃO DE NIL ═══")

	// Exemplo prático: função que retorna ponteiro ou nil
	var resultado *int = encontrarValor(25)

	if resultado != nil {
		fmt.Println("Valor encontrado:", *resultado) // Executa se encontrou
	} else {
		fmt.Println("Valor não encontrado") // Executa se retornou nil
	}

	// Com interface vazia: pode ser nil
	var dadosInterface interface{} = verificarDado(false)

	if dadosInterface != nil {
		fmt.Println("Dados disponíveis:", dadosInterface)
	} else {
		fmt.Println("Dados não disponíveis")
	}

	// ═══════════════════════════════════════════════════════════════════════════
	// 12. TYPE ASSERTION COM VERIFICAÇÃO - Alternativa segura
	// ═══════════════════════════════════════════════════════════════════════════

	fmt.Println("\n═══ 12. TYPE ASSERTION COM VERIFICAÇÃO ═══")

	// Sintaxe: valor, ok := variável.(tipo)
	// ok é true se a conversão foi bem-sucedida, false caso contrário

	var dados interface{} = "Hello"

	if str, ok := dados.(string); ok {
		fmt.Printf("Conversão bem-sucedida: %s\n", str) // Executa
	} else {
		fmt.Println("Não é uma string")
	}

	// ═══════════════════════════════════════════════════════════════════════════
	// RESUMO DE BOAS PRÁTICAS
	// ═══════════════════════════════════════════════════════════════════════════

	/*
			USE if/else QUANDO:
			- Apenas algumas condições possíveis
			- Condições complexas com &&, ||
			- Teste lógico simples

			USE switch QUANDO:
			- Muitas opções possíveis
			- Comparando um valor contra múltiplas constantes
			- Código mais legível e eficiente

			USE switch sem argumento QUANDO:
			- Condições complexas diferentes em cada branch
			- Quer escrever if/else de forma mais organizada

		LEMBRE-SE:
		- Em Go, não há operador ternário (a ? b : c)
		- Use if/else em uma linha se quiser: if x { y } else { z }
		- Break é automático em switch (não precisa escrever)
		- default é opcional tanto em if quanto em switch
		- Sempre verifique se ponteiros são nil antes de usar
	*/

}

// ═══════════════════════════════════════════════════════════════════════════
// FUNÇÕES AUXILIARES PARA EXEMPLOS
// ═══════════════════════════════════════════════════════════════════════════

// Função que retorna um ponteiro ou nil
// Utilizada no exemplo de verificação de nil
func encontrarValor(numero int) *int {
	if numero > 0 {
		return &numero
	}
	return nil // Retorna nil se a condição não for atendida
}

// Função que retorna interface{} ou nil
// Utilizada no exemplo de verificação de nil
func verificarDado(condicao bool) interface{} {
	if condicao {
		return "Dado disponível"
	}
	return nil // Retorna nil se a condição não for atendida
}
