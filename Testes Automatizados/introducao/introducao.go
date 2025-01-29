package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	
	endereco := "Avenida paulista"
	tipo := enderecos.TipoEnderecos(endereco)
	fmt.Println(tipo)
}