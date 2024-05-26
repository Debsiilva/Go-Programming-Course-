package main

import "fmt"

//a linguagem go não é orientada a objetos
//struct é como se fosse uma estrutura onde você armazena dados
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}
	wesley.Ativo = false

	fmt.Printf("Nome : %s, Idade: %d, Ativo : %t", wesley.Nome, wesley.Idade, wesley.Ativo)

}
