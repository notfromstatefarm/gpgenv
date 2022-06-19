package cmd

import (
	"fmt"
	"github.com/notfromstatefarm/gpgenv/internal/store"
	"io/ioutil"
	"os"
	"os/exec"
)

func Edit() {
	f, err := ioutil.TempFile(os.TempDir(), "gpgenv")
	if err != nil {
		fmt.Printf("gpgenv: failed to create temporary file: %v\n", err)
		os.Exit(1)
	}

	if store.Exists() {
		rawStore, err := store.ReadRaw()
		if err != nil {
			fmt.Printf("gpgenv: failed to read encrypted store: %v\n", err)
			os.Exit(1)
		}

		_, err = f.Write(rawStore)
		if err != nil {
			fmt.Printf("gpgenv: failed to write temporary file: %v\n", err)
			os.Exit(1)
		}
	}

	preStat, err := f.Stat()
	if err != nil {
		fmt.Printf("gpgenv: failed to stat temporary file: %v\n", err)
		os.Exit(1)
	}

	err = f.Close()
	if err != nil {
		fmt.Printf("gpgenv: failed to close temporary file: %v\n", err)
		os.Exit(1)
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}

	cmd := exec.Command(editor, f.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("gpgenv: failed to run editor: %v\n", err)
		os.Exit(1)
	}

	postStat, err := os.Stat(f.Name())
	if err != nil {
		fmt.Printf("gpgenv: failed to stat temporary file: %v\n", err)
		os.Exit(1)
	}

	if preStat.ModTime() == postStat.ModTime() {
		fmt.Printf("gpgenv: no changes made\n")
		return
	} else {
		newData, err := ioutil.ReadFile(f.Name())
		if err != nil {
			fmt.Printf("gpgenv: failed to read temporary file: %v\n", err)
			os.Exit(1)
		}
		err = os.Remove(f.Name())
		if err != nil {
			fmt.Printf("gpgenv: failed to remove temporary file: %v\n", err)
			os.Exit(1)
		}
		s, err := store.Unmarshal(newData)
		if err != nil {
			fmt.Printf("gpgenv: changes invalid: %v\n", err)
			os.Exit(1)
			return
		}
		err = s.Write()
		if err != nil {
			fmt.Printf("gpgenv: failed to write encrypted store: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("gpgenv: store written to disk\n")
	}
}
