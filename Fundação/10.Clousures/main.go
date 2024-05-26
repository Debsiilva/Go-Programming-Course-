package main

import "fmt"

func main() {
	//função clousures que seria uma função dentro de uma função
	total := func() int {
		return sum(1, 3, 45, 6, 34, 654, 7645, 534, 543, 543, 543) * 2
	}()

	fmt.Println(total)

}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
