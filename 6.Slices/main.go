package main

import "fmt"

type ID uint16

var (
	e float64 = 1.2
	f ID      = 1
)

func main() {
	s := []int{10, 20, 30, 50, 60, 70, 80, 90, 100}     // slice é declarado quando o [] estiver vazio , o slice não é declarado o tamanho dele
	fmt.Printf("len=%d cap %d %v\n", len(s), cap(s), s) // tamanho e capacidade e o valores
	// : criar um corte ( cria um slice de um slice)
	fmt.Printf("len=%d cap %d %v\n", len(s[:0]), cap(s[:0]), s[:0]) //:0 (direita) quer dizer que no começo do slice até o final será apagado

	fmt.Printf("len=%d cap %d %v\n", len(s[:4]), cap(s[:4]), s[:4]) //:4 (direita) quer dizer que os 4 primeiros vaçpres do slice será mantido

	fmt.Printf("len=%d cap %d %v\n", len(s[2:]), cap(s[2:]), s[2:]) //2: (esquerda) quer dizer que os 2 primeiros valores do slice será apagado

	s = append(s, 110) // o append irá duplicar a capacidade do array inicial se ele não tiver a capacidade suficiente

	fmt.Printf("len=%d cap %d %v\n", len(s[:2]), cap(s[:2]), s[:2])

	fmt.Printf("len=%d cap %d %v\n", len(s), cap(s), s)

}
