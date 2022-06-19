package main

import (
	"fmt"
	"github.com/notfromstatefarm/gpgenv/internal/cmd"
	"os"
)

func printUsage() {
	fmt.Printf(`Executes commands with environment variables stored at rest via gpg

Usage:
  gpgenv [context] [command]		Executes a command
  gpgenv edit				Opens your editor to edit .gpgenv (set EDITOR environment variable to change editor)
`)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "help" || args[0] == "--help" {
		printUsage()
		return
	}

	switch args[0] {
	case "edit":
		cmd.Edit()
	default:
		cmd.Run()
	}

}
