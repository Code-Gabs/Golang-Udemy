// TESTE UNITÁRIO
package enderecos

import (
	//"strings"
	"testing"
)

// a struct é responsável por colocar os campos que precisam usar nos testes como forma de paramentros
type cenarioTeste struct {
	enderecoRecebido string
	retornoEsperado string
}

// ao criar um metodo q comece com Test o go entende q é um funcção de testes
func TestTiposEnderecos(t *testing.T) {

	CenariosTests := []cenarioTeste {
		
		{"Rua jardins", "RUA"},
		{"avenida dos bobos", "AVENIDA"},
		{"estrada dos galdinos", "ESTRADA"},
		{"Rodovia Rodoanel", "RODOVIA"},
		{"Praça eliseu junior", "PRAÇA"},
		{"PRAÇA DA ESPERANÇA", "PRAÇA"},
		{"", "Endereço invalido - Tipo não cadastrado"},
	}

	for _, cenario := range CenariosTests {
		stringRecebida := TipoEnderecos(cenario.enderecoRecebido)
		if stringRecebida != cenario.retornoEsperado {
			t.Errorf("O tipo recebido: %s é diferente do esperado: %s", stringRecebida, cenario.retornoEsperado)
		}
	}
}