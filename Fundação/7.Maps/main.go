package main

import "fmt"

func main() {

	salarios := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	//fmt.Println(salarios["Wesley"])
	delete(salarios, "Wesley")
	salarios["Wes"] = 5000
	//fmt.Println(salarios["Wesley"])

	/* 	sal := make(map[string]int) // função make ele prepara um map mas vazio
	   	sal1 := map[string]int{}    // cria um map vazio também
	   	sal1["Wesley"] = 1000 */

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d \n", nome, salario)
	}
	for _, salario := range salarios { // _ é o blank identifier ele vai ignorar o valor esperado para estar ali
		fmt.Printf("O salario é %d \n", salario)
	}

}
