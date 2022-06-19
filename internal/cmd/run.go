package cmd

import (
	"fmt"
	"github.com/notfromstatefarm/gpgenv/internal/store"
	"os"
	"os/exec"
)

func Run() {
	contextName := os.Args[1]
	command := os.Args[2]
	args := os.Args[3:]

	if !store.Exists() {
		fmt.Printf("gpgenv: store not configured, run gpgenv edit")
		os.Exit(1)
	}

	store, err := store.Read()
	if err != nil {
		fmt.Printf("gpgenv: could not read store: %v", err)
		os.Exit(1)
	}

	context, ok := store.Contexts[contextName]
	if !ok {
		fmt.Printf("gpgenv: context %s not defined in store", contextName)
	}

	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	for k, v := range context {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		} else {
			fmt.Printf("gpgenv: could not run command: %v", err)
			os.Exit(1)
		}
	} else {
		os.Exit(0)
	}
}
