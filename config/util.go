package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func path(p string) string {
	if strings.HasPrefix(p, "/") {
		return p
	}

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if strings.HasPrefix(p, "./") {
		p = strings.Replace(p, "./", "", -1)
	}

	return filepath.Join(path, p)
}
