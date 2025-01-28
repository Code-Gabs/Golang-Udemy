package main


import "fmt"

func main() {
	
	// PONTEIROS SÃO REFERENCIAS DE MEMÓRIA DE ALGUMA INFORMAÇÃO QUE FOI ARMAZENADA OU CRIADA
	// O objetivo de usar Ponteiros é de utilizar uma informação que está em memoria sem precisar criar cópias, dessa forma podemos reutilizar um dado sempre que quisermos e modificar o valor dele
	var numero int
	var ponteiro *int //Para identificar como ponteito basta colocar "*" na frente do tipo, um ponteito que não possui valor é igual a NULL

	fmt.Println(numero)
	fmt.Println(ponteiro)

	numero = 100

	// REFERENCIA DE MEMORIA DA VARIAVEL "numero"
	ponteiro = &numero // com o identificador "&" voce consegue referenciar uma informação de alguma variavel

	fmt.Println(numero)
	fmt.Println(ponteiro)
	fmt.Println("Mostrar o valor da referencia de memoria do Ponteiro: ", *ponteiro)

}