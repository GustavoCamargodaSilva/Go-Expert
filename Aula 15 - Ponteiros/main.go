package main

type Cliente struct {
	nome  string
	idade int
}

func soma(a *int) {
	*a = 100
}

func main() {

	//Memória  -> Endereço -> Valor ->
	// a := 10
	// fmt.Println(a)

	// soma(&a)
	// fmt.Println(a)

}
