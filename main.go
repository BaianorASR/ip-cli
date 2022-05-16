package main

import (
	"log"
	"modulo/app"
	"os"
)

func main() {
	println("Initializing...")

	app := app.Generate()
	if erro := app.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
