package main

import "fmt"

type ID uint16

var (
	e float64 = 1.2
	f ID      = 1
)

func main() {
	var meuArray [3]int // o array no momento que ele é declarado ele tem um tamanho fixo
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 10
	fmt.Println(len(meuArray) - 1)         // descobrir o ultimo elemento de um array
	fmt.Println(meuArray[len(meuArray)-1]) // descobrir o valor do ultimo elemento

	for i, v := range meuArray { //range é para percorrer todos os itens do array
		fmt.Printf("O indice é %d e o valor é %d\n", i, v) // se for string coloca v se for digito coloca d
	}

}
