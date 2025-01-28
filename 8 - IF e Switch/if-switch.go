package main


import "fmt"

func dayOfWeekend(dia int) string{

	switch dia {
	case 1:
		return "Segunda-feira"
	case 2:
		return "Terça-feira"
	case 3:
		return "Quarta-feira"
	case 4:
		return "Quinta-feira"
	case 5:
		return "Sexta-feira"
	case 6:
		return "Sabado"	
	case 7:
		return "Domingo"
	default:
		return "Número invalido"
	}

	
}


func main() {
	
	dia := dayOfWeekend(4)
	fmt.Println(dia)
}