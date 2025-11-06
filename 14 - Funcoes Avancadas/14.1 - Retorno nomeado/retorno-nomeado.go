package main

func calculosMatematios(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculosMatematios(10, 5)
	println("Soma:", soma)
	println("Subtração:", subtracao)
}
