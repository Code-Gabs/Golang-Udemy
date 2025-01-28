package main

import "fmt"

func RecuperarExecution() {
	if r:= recover(); r != nil {
		fmt.Println("Tentativa de recuperar execução com sucesso!")
	}
}

func CalculoAlunoNota(n1, n2 float64) bool {
	media := (n1 + n2) / 2
	defer RecuperarExecution()
	
	if media < 6 {
		return true
	} else if media > 6 {
		return false
	}

	panic("\nA MÉDIA DO ALUNO É IGUAL A 6 !!!")


}

func main() {
	fmt.Println(CalculoAlunoNota(4,0))
	
	fmt.Println("EXECUTANDO DEPOIS DO PANIC")
}