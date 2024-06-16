package main

func main() {

	//trabalhando com ponteiros , o ponteiro nada mais é do que um endereço na memória que tem um valor , as váriaveis apontam para um endereço na memória que tem um valor
	// se você pegar um ponteiro que é esse endereço na memória e mudar o valor quando essa variável acessar pode ter mudado o valor
	// é usado bastante na linguagem , é utilizado para ser utilizada não só em um escopo local mais em qualquer local

	// Memória -> Endereço -> Valor
	//Guardando uma caixinha que tem um endereço e que tem o valor 10
	// variavel -> ponteiro que tem um endereço na memória -> valor
	a := 10
	//qual é a caixa do endereço que você está guardando
	println(&a)
	// o * diz respeito ao endereço de a que seria como mostrado no terminal o 0xc000059f38
	// o ponteiro é o endereçamento na memória
	var ponteiro *int = &a

	println(ponteiro)
	// agora o endereço da memória valera 20
	*ponteiro = 20
	println(a)

	// neste caso o b está sendo um ponteiro , ele está apontando para o mesmo endereço de memória de a
	b := &a
	println(b)

	//como sei o valor que o b está apontando
	//basta coloca o * = reference , qual é o valor guardado na memória
	println(*b)

}
