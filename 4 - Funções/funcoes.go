package main

// o int após os paramêtros indica o tipo do valor retornado pela função
func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	println(soma)

	var f = func(txt string) string {
		println(txt)
		return txt
	}

	resultado := f("Função como variável")
	println(resultado)

	// resultadoSoma, resultadoSubtracao := calculosMatematicos(15, 10)
	// println(resultadoSoma)
	// println(resultadoSubtracao)

	// caso não queira usar as duas variáveis retornadas pela função, pode usar o underline _ para ignorar uma delas
	resultadoSoma, _ := calculosMatematicos(15, 10)
	println(resultadoSoma)

}
