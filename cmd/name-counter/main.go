package main

import (
	"os"

	"name-counter/cmd/name-counter/internal/cli"
)

func main() {
	if err := cli.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
