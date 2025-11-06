package main

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	soma1 := soma(1, 2, 3)
	soma2 := soma(10, 20, 30, 40, 50)

	println("Soma 1:", soma1)
	println("Soma 2:", soma2)
}
