package main

import (
	"github.com/urfave/cli"
	"log"
)

func main() {
	log.Printf("vayne attack")
	app := cli.NewApp()
	app.Name = "vayne"
	app.Usage = "ATE- AutoTest"
	app.Version = "1.0.0"

}
