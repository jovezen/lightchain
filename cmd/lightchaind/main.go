package main

import (
	"os"

	"github.com/jovezen/lightchain/cmd/lightchaind/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
