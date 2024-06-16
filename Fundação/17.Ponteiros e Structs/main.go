package main

/* type Cliente struct {
	nome string
} */
type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

/* func (c Cliente) andou() {
	c.nome = "Wesley Willians"
	fmt.Printf("O cliente %v andou\n ", c.nome)
} */
/* func (c Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
} */
func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

/* func main() {
	Wesley := Cliente{
		nome: "Wesley",
	}

	Wesley.andou()
	fmt.Printf("O valor da struct com nome %v", Wesley.nome)
} */
func main() {
	conta := Conta{saldo: 100}
	conta.simular(200)
	println(conta.saldo)
}
