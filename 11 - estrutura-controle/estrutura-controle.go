package main

import "fmt"

func main() {
	numero := 10

	if numero > 15 {
		println("Número maior que 15")
	} else {
		println("Número menor que 5")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("numero maior que 0")
	} else {
		fmt.Println("numero menor ou igual a 0")
	}
}
