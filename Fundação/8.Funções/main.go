package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(50, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

/* func sum(a int, b int) int {
	return a + b
}
*/
//quando são do mesmo tipo da para simplificar desta forma
/* func sum(a, b int) int {
	return a + b
} */

// Em Go a função pode retornar dois tipos de tipagens diferentes
/* func sum(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}
	return a + b, false
} */

// como o Go não tem as exceptions que nem em outras linguagens , normalmente é trabalhado dessa forma as funções
// para identificação de erros
func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
