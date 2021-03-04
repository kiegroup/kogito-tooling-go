package utils

import (
	"os"
	"path/filepath"
)

func GetBaseDir() string {
	return filepath.Dir(os.Args[0])
}
