package main

import (
	"discover_server/app"
	"log"
	"os"
)

func main() {

	aplication := app.App()

	if err := aplication.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
