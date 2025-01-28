package main
import (
	"fmt"
	"time"
	)

	
func alwaysPrint(texto string, chanel chan string) {

	for i := 0; i <= 5 ; i++{
		// aqui informamos qual a informação será enviada para o canal
		// a ordem da seta altera o fluxo de envio e recebimento
		// da forma como esta abaixo, estamos enviando uma string para o canal
		chanel <- texto
		time.Sleep(time.Second)
	}

	// instrução para que feche o canal e assim ele não irá aguardar novas mensagens
	// evita de apresentar erro - DeadLock
	close(chanel)
}

func main() {
	// criando um canal em Go
	// ao criar um canal deve ser passado o tipo de dado
	canal := make(chan string)
	go alwaysPrint("Ola mundo", canal)

	// Dessa forma estamos informado que o canal irá receber uma informação
	// ele irá esperar essa informação chegar para fazer ações, como se estivesse dando uma ordem
	for {
		// podemos verificar o estado do canal, aberto(true) ou fechado(false)
		mensage, aberto := <- canal
		if !aberto {
			break
		}
		fmt.Println(mensage, aberto)
		
	}
	
	// outra forma de usar o for acima
	// for mensagens := range canal {
		
	// 	fmt.Println(mensagens)
	// }
	
	fmt.Println("Fim da execução")
	
}