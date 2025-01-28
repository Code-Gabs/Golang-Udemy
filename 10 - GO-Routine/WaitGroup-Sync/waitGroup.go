package main

import (
	"fmt"
	"sync"
	"time"
)


func alwaysPrint(texto string) {

	for i := 0; i <= 5 ; i++{
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {

	// criando um waitGroup
	var waitGroup sync.WaitGroup

	// adiciono quantas rotinas serão aguardadas
	waitGroup.Add(2)

	go func ()  {
		alwaysPrint("Ola mundo GOROUTINE")
		waitGroup.Done() // quando colocamos o Done, mostra para o Go que ele tira da fila de execução "-1" no contador adicionado na linha 23
		
	}()

	go func ()  {
		alwaysPrint("Continuando execução")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}