package main

import (
	"fmt"
)

func main() {
	// i := 0
	// for i < 10 {
	// 	i++
	// 	fmt.Println("incrementando i")
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)

	// for j := 0; j < 10; j += 2 {
	// 	fmt.Println("incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Ana", "Bia", "Carlos"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	usuario := map[string]string{
		"nome":      "João",
		"sobreNome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
