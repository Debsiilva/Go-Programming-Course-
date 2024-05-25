package main

import "fmt" // pacote format ele auxilia na formatação dos dados ao executa-lo

type ID uint16 // criando tipagem em GO

var (
	e float64 = 1.2
	f ID      = 1
)

func main() {
	fmt.Printf("O tipo de E é %T", e)  //%T retorna o tipo da variável que você está mencionando
	fmt.Printf("O valor de E é %v", e) //%v retorna o valor da variável que você está mencionando
	fmt.Printf("O tipode F é %T", f)
}
