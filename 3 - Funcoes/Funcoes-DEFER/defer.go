package main

import "fmt"

func funcao1 () {
	fmt.Print("Executando funcao 1")
}


func funcao2 () {
	fmt.Print("Executando funcao 2")
}

// Muito usado no contexto de conexão com o banco de dados
// onde sempre que uma conexao é aberta é preciso fechar
func main() {
	
	funcao2()
	// DEFER adiar uma função
	defer funcao1()
}