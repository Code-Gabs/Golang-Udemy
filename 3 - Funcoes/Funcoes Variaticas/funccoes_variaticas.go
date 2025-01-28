package main

import "fmt"

func soma(num ...int) (result int) {

	soma := 0

	for _, n := range num {
		fmt.Println("Numero para somar\n", n)
		soma += n
	}
	result = soma

	return
}

// Funções variaticas devem receber apena 1 parametro desse tipo e sempre o ultimo paramentro da função
func mostrarTexto(text string, numeros ...int) {
	for _, index := range numeros {
		fmt.Println(text, index)
	}
}

func main() {
	result := soma(1,2,3,89,478,378)
	fmt.Println(result)

	mostrarTexto("Comprando...", 12,44,6565,76,454,43)
}