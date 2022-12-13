package main

import (
	"log"

	"github.com/kubekumber/kubekumber/cmd"
)

var version string = "v0.0.0"

func main() {
	if err := cmd.Execute(version); err != nil {
		log.Fatal(err)
	}
}
