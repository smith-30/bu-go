package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
	"github.com/smith-30/bu-go/helper/env"
	"github.com/smith-30/bu-go/subcmd"
)

var (
	revision = "unknown"
)

func main() {
	env.LoadEnv()

	c := cli.NewCLI("bu-go", revision)

	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		subcmd.UPLOAD: func() (cli.Command, error) {
			return &subcmd.Upload{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
