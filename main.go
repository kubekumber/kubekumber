package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/bitfield/script"
	cli "github.com/urfave/cli/v2"
)

func main() {

	var command string
	var regex string
	var verbose bool
	var dry_run bool

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "regex",
			Aliases:     []string{"r"},
			Value:       getCurrentConfig(),
			Usage:       "regex for cluster selection",
			Destination: &regex,
		},
		&cli.StringFlag{
			Name:        "command",
			Aliases:     []string{"c"},
			Value:       "",
			Usage:       "command to run on cluster",
			Destination: &command,
		},
		&cli.BoolFlag{
			Name:        "verbose,v",
			Aliases:     []string{"v"},
			Usage:       "print cluster name and command before every output",
			Destination: &verbose,
		},
		&cli.BoolFlag{
			Name:        "dry-run,d",
			Aliases:     []string{"d"},
			Usage:       "Print out clusters and command that would be run",
			Destination: &dry_run,
		},
	}

	app.Action = func(c *cli.Context) error {

		var clusters []string

		if regex == getCurrentConfig() {

			clusters = []string{regex}

		} else {

			regex_compiled := regexp.MustCompile(regex)

			clusters_rune, _ := script.Exec("kubectx").MatchRegexp(regex_compiled).String()

			clusters = strings.Split(strings.TrimSpace(clusters_rune), "\n")

		}

		for _, cluster := range clusters {

			if verbose || dry_run {
				fmt.Println("[ DEBUG ] CONTEXT: " + cluster)
				fmt.Println("[ DEBUG ] COMMAND: kubectl " + command)
			}

			if !dry_run {
				script.Exec("kubectl " + command + " --context " + cluster).Stdout()
			}

		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getCurrentConfig() string {
	cluster, _ := script.Exec("kubectl config current-context").String()
	return strings.TrimSpace(cluster)
}
