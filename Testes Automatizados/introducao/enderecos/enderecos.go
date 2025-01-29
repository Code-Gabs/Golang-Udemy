package enderecos

import (
	"fmt"
	"strings"
)

func TipoEnderecos(endereco string) string {
	
	tiposEnderecosValidos := []string{"rua", "praça", "rodovia", "avenida", "estrada"}
	enderecoMinusculo := strings.ToLower(endereco)

	primeiraPalavraEndereco := strings.Split(enderecoMinusculo, " ")[0]

	enderecoValido := false
	for _, tipos := range tiposEnderecosValidos {
		if tipos == primeiraPalavraEndereco {
			enderecoValido = true
			fmt.Println("Endereço válido")
		}
	}

	if enderecoValido {
		return strings.ToUpper(primeiraPalavraEndereco)
	}

	return "Endereço invalido - Tipo não cadastrado"
}