package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Maria",
			"ultimo":   "Souza",
		},
		"curso": {
			"nome":  "Engenharia",
			"nivel": "Bacharelado",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["idade"] = map[string]string{
		"anos": "23",
	}
	fmt.Println(usuario2)
}
