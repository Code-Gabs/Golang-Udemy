package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main()  {
	fmt.Println("HELLOU WORLD")
	auxiliar.WriteText()

	// essa função valida o formato do email informado, podendo ser printada no console
	erro := checkmail.ValidateFormat("go@@email.com")
	fmt.Println(erro)
}