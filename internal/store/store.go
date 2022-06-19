package store

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"gopkg.in/yaml.v2"
)

type Store struct {
	Contexts map[string]map[string]string `yaml:"contexts"`
	KeyEmail string `yaml:"key-email"`
}

func getPath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", dir, ".gpgenv"), nil
}

func Exists() bool {
	path, err := getPath()
	if err != nil {
		return false
	}
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}

func ReadRaw() ([]byte, error) {
	if !Exists() {
		return nil, errors.New("store doesn't exist")
	}
	path, err := getPath()
	if err != nil {
		return nil, err
	}
	cmd := exec.Command("gpg", "--decrypt", path)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}

func Read() (*Store, error) {
	raw, err := ReadRaw()
	if err != nil {
		return nil, err
	}
	var store Store
	err = yaml.Unmarshal(raw, &store)
	if err != nil {
		return nil, err
	}
	return &store, nil
}

func Unmarshal(data []byte) (*Store, error) {
	var store Store
	err := yaml.Unmarshal(data, &store)
	if err != nil {
		return nil, err
	}
	return &store, nil
}

func (s *Store) ToRaw() ([]byte, error) {
	return yaml.Marshal(s)
}

func (s *Store) Write() error {
	path, err := getPath()
	if err != nil {
		return err
	}

	data, err := s.ToRaw()

	buffer := bytes.Buffer{}
	buffer.Write(data)

	cmd := exec.Command("gpg", "--encrypt", "--yes", "-r", s.KeyEmail, "-o", path)
	cmd.Stdin = &buffer
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	err = os.Chmod(path, 0600)
	if err != nil {
		return err
	}
	
	return nil
}