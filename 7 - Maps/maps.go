package main

import "fmt"

func main() {
	
	// como criar um map -> Um map deve possuir sua chave e seu valor com o mesmo tipo (Limitação para maps diferente de struct)

	aluno := map[string]string {
		"nome": "Gabriel",
		"idade": "25",
	}

	fmt.Println(aluno)
	fmt.Println(aluno["nome"])

	funcionario := map[string]map[string]string {
		"info" : {
			"nome": "Gabriel",
			"matri_id": "189231",
			"cargo": "Analista JR",
		},

		"dados": {
			"signo": "Capricornio",
		},
	}

	fmt.Println(funcionario)
	
	delete(funcionario, "dados")
	fmt.Println(funcionario)
	
	// adicionar valores a um map
	funcionario["salario"] = map[string]string{
		"valor": "7500",
	}
	fmt.Println(funcionario)
}