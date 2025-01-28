package main

import "fmt"

type usuario struct {
	nome	string
	id	float64
	idade 	int
	active	bool
}

// em go não existe Orientaado a Objeto
// como criar um metodo em que seja somente para essa "classe" ou obejto
// a função abaixo referencia um struct do tipo "usuario" criado anteriormente e possui um nome e pode conter um return também
func (u usuario) mostrarDados() {
	fmt.Printf("Dados de dentro do usuario %s que está ativo = %v\n", u.nome, u.active)
}


// função para fazer um update nos dados de um structu usando metodo
// ao realizar um update de dados é importante usar os Ponteiros dessa formsa evita a cópia de valores que precisam ser atualizados de fato
// caso não use um Ponteiro a forma como o dado ficará pode afetar em não atualizar de fato um campo no banco, criando cópias desatualizadas
func (u *usuario) happyBirthday() {
	u.idade++
}


func main() {
	user := usuario{"Gabriel", 483993,23, true}
	fmt.Println(user)
	user.mostrarDados()
	user.happyBirthday()
	fmt.Printf("Aniversariante do dia %v, Parabens %s ", user.idade, user.nome)
}