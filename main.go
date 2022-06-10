package main

import (
	"busca-ip-servers/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting....")

	aplicacao := app.Gerar()
	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
