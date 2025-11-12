package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a funcao 1")
}

func funcao2() {
	fmt.Println("Executando a funcao 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada com sucesso. Retornando resultado...")
	fmt.Println("Entrando na funcao para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	//DEFER = ADIAR
	fmt.Println("Aluno aprovado:", alunoAprovado(7, 8))
}
