// cmd/superdata/main.go
package main

import (
	"flag"
	"log"
	"os"

	"superdata/internal/superdata"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := superdata.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
