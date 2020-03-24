package utils

import (
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

// FileExists checks via os.IsNotExist of the file
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// WalkPath gathers all files in a given path
func WalkPath(walkPath string) []string {
	var files []string
	err := filepath.Walk(walkPath, func(path string, info os.FileInfo, err error) error {
		name := info.Name()
		if strings.Contains(name, ".json") {
			files = append(files, name)
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return files
}
