package main

import (
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "myCLI"
	app.Usage = "To get things and stuff done"
	app.Commands = []cli.Command{}

	app.Run(os.Args)
}
