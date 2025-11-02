package main

import "fmt"

type pessoa struct {
	nome      string
	sobreNome string
	idade     uint
	altura    uint
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"Gabriel", "Gallo", 23, 180}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "Univille"}
	fmt.Println(e1.nome)
}
