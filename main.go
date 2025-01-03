package main

import (
	"fmt"
	"os"

	"github.com/pammalPrasanna/idi/cmd/cli"
)

func main() {
	if err := cli.Main(); err != nil {
		fmt.Printf("idi: %s\n", err.Error())
		os.Exit(1)
	}
}
