package utils

import (
	"fmt"
	"os"
)

func CreateFile(name string, content []byte) error {
	if _, err := os.Stat(name); err == nil {
		return fmt.Errorf("such file exists %s", name)
	}
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Write(content); err != nil {
		return err
	}
	if err := file.Sync(); err != nil {
		return err
	}
	return file.Close()
}

func IsStatFile(name string) bool {
	if _, err := os.Stat(name); err == nil {
		return true
	}
	return false
}
