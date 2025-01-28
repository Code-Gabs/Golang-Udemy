package main

import "fmt"

type users struct {
	nome	string
	idade 	uint8
	profissao 	string
}

func main() {
	
	// 1° jeito de instanciar uma strucut e manipular os valores
	var user users

	user.nome = "Gabriel"
	user.idade = 24
	user.profissao = "QA Analyst"

	fmt.Println(user)

	// 2° jeito de instanciar uma strucut e manipular os valores
	user2 := users{
		"Felipe Valdo", 46, "CTO"}
	fmt.Println(user2)
	
	// 3° jeito de instanciar uma strucut e informar os valores para objetod especificos
	user3 := users{nome: "Rodolfo", profissao: "DEV -Backend"}
	fmt.Println(user3)
}