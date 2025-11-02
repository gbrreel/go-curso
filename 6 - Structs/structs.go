package main

type usuario struct {
	nome  string
	idade uint8
}

func main() {

	var u usuario
	u.nome = "João"
	u.idade = 25

	println(u.nome)
	println(u.idade)
}
