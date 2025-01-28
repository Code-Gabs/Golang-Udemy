package main

import (
	app "linha-de-comando/App"
	"log"
	"os"
)

func main() {

	application := app.Gerar()
	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}