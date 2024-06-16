package main

//para eu passar o endereço na momória eu terei que referência por ponteiro *
//func soma(a, b int) int {
// para eu passar os valores reais sem ser mais a cópia eu terei referênciar pelo ponteiro
func soma(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	println(&minhaVar2)
	//soma(minhaVar1, minhaVar2)
	soma(&minhaVar1, &minhaVar2)
	// aqui eu fiz uma cópia do valor . eu passei apenas o valor para a variável "a" la da função soma
	println(minhaVar1)
	println(minhaVar2)

}
