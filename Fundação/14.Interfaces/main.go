package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

//a interface está chamando o método desativar , ou seja, qualquer struct que criar que tenha o método desativar está chamando está interface
type Pessoa interface {
	//a interface do Go só permite declarar métodos
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

//compondo uma struct encadeando na outra
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

//método Desativar
func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado ", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {

	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}
	minhaEmpresa := Empresa{}
	Desativacao(minhaEmpresa)
	Desativacao(wesley)

}
