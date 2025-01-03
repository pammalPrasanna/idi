package main

import (
	"fmt"
	"os"
)

func main() {
	if err := Main(); err != nil {
		fmt.Printf("idi: %s\n", err.Error())
		os.Exit(1)
	}
}
