package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 3, 45, 6, 34, 654, 7645, 534, 543, 543, 543))
}

// nesse exemplo de função variádica é passado como parametro a variavel números e ... que significa
// que pode passar um número indeterminado de parametros
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
