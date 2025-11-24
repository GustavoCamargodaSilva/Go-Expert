package matematica

func Soma[T int | float64](a, b T) T {
	return a + b
}

// var a int = 10

type Carro struct { //Se a struct for com letra maiúscula mas
	Modelo string // -> o atributo for com letra minuscula, a struct fica visivel mas o Modelo nao.
}
