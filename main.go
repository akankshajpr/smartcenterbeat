package main

import (
	"os"

	"github.com/Microland/smartcenterbeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
