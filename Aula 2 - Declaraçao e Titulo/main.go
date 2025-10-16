package main

const a = "Hello, Go!"

var (
	b bool
	c int
	d string
	e float64

	//variaveis de escopo global na qual podemos acessar dentro de qualquer outra classe dentro do mesmo pacote
)

func main() {
	var a string = "X" //variavel privada ao método que fizemos a declaraçao

	f := "X" //forma mais fácil de declarar variável, só é feito na primeira vez que a variável é criada

	//f:= "fd"  -> nao posso fazer isso pq só podemos declarar assim na primeira vez

	println(a, f)
}
