package main

import (
	"application-line-command/app"
	"log"
	"os"
)

func main() {

	application := app.Generate()

	er := application.Run(os.Args)
	if er != nil {
		log.Fatal(er)
	}
}
