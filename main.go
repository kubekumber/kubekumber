package main

import (
	"github.com/kubekumber/kubekumber/cmd"
)

var version string = "v0.0.0"

func main() {
	cmd.Execute(version)
}
