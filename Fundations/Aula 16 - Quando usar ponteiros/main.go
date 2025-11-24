package main

// COM PONTEIRO (Mutável)
// Esta função usa ponteiros como parâmetros
// Isso permite MODIFICAR os valores originais das variáveis
//
// Quando usar ponteiros:
// 1. Quando você precisa alterar o valor de uma variável dentro de uma função
// 2. Quando quer evitar copiar grandes estruturas de dados (structs grandes)
// 3. Quando precisa compartilhar a mesma memória entre múltiplas partes do código
//
// Desvantagem: Mais complexo de entender e pode causar efeitos colaterais inesperados
func soma(a, b *int) int {
	// Desreferência: *a acessa o valor que 'a' aponta
	// Isso muda o valor ORIGINAL de minhaVar1, não uma cópia
	*a = 50
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20

	// Passamos o endereço das variáveis usando &
	// Assim a função pode modificar os valores originais
	println(soma(&minhaVar1, &minhaVar2)) // Saída: 70 (50 + 20)

	// minhaVar1 foi modificada para 50 dentro da função soma
	println(minhaVar1) // Saída: 50

	teste()
}

// SEM PONTEIRO (Imutável - Passa por valor)
// Esta função recebe cópias dos valores, NÃO os originais
//
// Quando usar valores comuns (sem ponteiros):
// 1. Quando você NÃO precisa modificar a variável original
// 2. Quando quer garantir que a função não tenha efeitos colaterais
// 3. Quando trabalha com tipos pequenos (int, string, bool, etc)
// 4. Quando quer código mais seguro e previsível
//
// Vantagem: Mais seguro, sem efeitos colaterais inesperados
func soma1(a, b int) int {
	// Esta atribuição modifica apenas a CÓPIA local de 'a'
	// O valor original de minhaVar1 não é afetado
	a = 50
	return a + b
}

// Esta função demonstra a diferença entre usar ponteiros e valores
func teste() {
	minhaVar1 := 10
	minhaVar2 := 20

	// Passamos os VALORES diretos (sem &)
	println(soma1(minhaVar1, minhaVar2)) // Saída: 70 (50 + 20)

	// minhaVar1 continua 10 porque a função modificou apenas uma cópia
	println(minhaVar1) // Saída: 10
}

// COMPARAÇÃO IMPORTANTE:
//
// Com Ponteiro (soma):
//   - Entrada: &minhaVar1 (endereço) e &minhaVar2 (endereço)
//   - Resultado: minhaVar1 muda para 50 ✓ Modificou original
//
// Sem Ponteiro (soma1):
//   - Entrada: minhaVar1 (valor 10) e minhaVar2 (valor 20)
//   - Resultado: minhaVar1 continua 10 ✗ Modificação local apenas
//
// REGRA PRÁTICA:
// - Use ponteiros quando: precisa modificar, trabalha com structs grandes
// - Use valores quando: quer segurança, dados são pequenos, sem modificações
