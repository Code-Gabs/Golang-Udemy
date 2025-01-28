package main

import (
	"fmt"
	"time"
)


func alwaysPrint(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
func main() {
	 // quando informamos "go" na frente de uma função, 
	 //significa que não será preciso espear o fim da execução dela para executar a proxima função
	go alwaysPrint("Ola mundo GOROUTINE")
	alwaysPrint("Continuando execução")
}