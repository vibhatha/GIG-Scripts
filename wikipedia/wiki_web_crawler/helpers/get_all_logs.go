package helpers

import (
	"log"
	"os"
	"path/filepath"
	"sort"
)

func getAllLogs(logPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(logPath, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Println("error loading log files")
		return []string{}, err
	}
	sort.Strings(files)
	return files, nil
}
