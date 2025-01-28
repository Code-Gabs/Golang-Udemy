package main

import "fmt"

func main() {
	// Array possuem tamanho e tipo fixo, não podendo alterar eles após a criação do Array
	// como criar um array: deve possuir o numero de posições e o tipo
	var lista [4]int

	lista[3] = 23

	fmt.Println(lista)

	// dessa foram vc nãop coloca um numero do tamanho, porem colocando o numero de itens entre {} isso será o tamanho do array
	array := [...]string{"Frtuas","Filmes", "Trabalho", "Dinheiro"}
	fmt.Println(array)
	
	fmt.Println("\nSlices")
	
	// Slices são mais flexiveis e possuem tamanho e tipos dinamicos, podendo fazer incrementos durante o uso 
	// Sileces não é necessário informar o tamanho
	slice1 := []int{1,34,3,4,69,67,76}
	fmt.Println(slice1)
	slice1 = append(slice1, 90)
	fmt.Println(slice1)

}