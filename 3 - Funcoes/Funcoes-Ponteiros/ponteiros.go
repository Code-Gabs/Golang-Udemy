package main

import "fmt"

func inverteSinalNumero(num int64) int64{
	return num * -1

}

func inverterSinalNUmeroPonteiros(num *int) {
	*num = *num * -1
}

func main() {
	numero := 10
	novoNumero := inverteSinalNumero(int64(numero))
	fmt.Println(numero)
	fmt.Println(novoNumero)

	/// funções com ponteiros
	numero1 := 33
	fmt.Println(numero1)
	inverterSinalNUmeroPonteiros(&numero1)
	fmt.Println(numero1)

}