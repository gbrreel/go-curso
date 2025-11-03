package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"

	default:
		return "Número inválido"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda-feira"
	case numero == 3:
		diaDaSemana = "Terça-feira"
	case numero == 4:
		diaDaSemana = "Quarta-feira"
	default:
		diaDaSemana = "Número inválido"
	}
	return diaDaSemana
}

func main() {
	dia := diaDaSemana(4)
	fmt.Println(dia)

	fmt.Println("--------")
	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)
}
