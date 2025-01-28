package main

import "fmt"

// função que pode retornar dois resultados diferentes
func calculoMath(n1,n2 int8) (int8, int8) {
	
	soma := n1 + n2
	subtration := n1 - n2
	return soma, subtration
}

// Função init() no go é uma função que sempre será executada ainda da função main()
// muito usada para setup de alguma coisa antes da execução principal
func init() {
	fmt.Println("Executando função init")
}


func returnText(text string) (string){
	return text
}

func main() {
	//resultadoSoma, resultadoSub := calculoMath(20,10)
	//fmt.Println(resultadoSoma, resultadoSub)

	// caso não queria usar os dois resultado fazer dessa forma declarando um "_"
	resultadoSoma, _ := calculoMath(20,10)
	fmt.Println(resultadoSoma)

	result := returnText("Gabriel")
	fmt.Println(result)
}