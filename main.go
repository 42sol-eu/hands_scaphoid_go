package main

import (
	"fmt"
	"os"

	"github.com/42sol-eu/hands_scaphoid_go/cmd_new"
)

func main() {
	if err := cmd_new.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
