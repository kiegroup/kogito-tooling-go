package utils

import (
	"os"
	"path/filepath"
)

func GetBaseDir() string {
	env := os.Getenv("ENV")
	if env != "dev" {
		return filepath.Dir(os.Args[0])
	} else {
		return "./"
	}
}
