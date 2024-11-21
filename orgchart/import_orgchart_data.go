package main

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const (
	fileExt = ".csv"
)

func main() {

	// Check if the root directory is provided as an argument
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <rootDir>", os.Args[0])
	}
	rootDir := os.Args[1]

	filePaths, err := getCSVFilePaths(rootDir)
	if err != nil {
		log.Fatalf("Error retrieving CSV file paths: %v", err)
	}

	sort.Slice(filePaths, func(i, j int) bool {
		date1 := extractDateFromFileName(filePaths[i])
		date2 := extractDateFromFileName(filePaths[j])
		return date1.Before(date2)
	})

	for _, filePath := range filePaths {
		unixPath := strings.ReplaceAll(filePath, `\`, "/")
		import_csv(unixPath)
	}
}

func getCSVFilePaths(dir string) ([]string, error) {
	var filePaths []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), fileExt) {
			filePaths = append(filePaths, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return filePaths, nil
}

func extractDateFromFileName(filePath string) time.Time {
	fileName := filepath.Base(filePath)
	gazetteDate, err := time.Parse("gazette-2006-1-2.csv", fileName)
	if err != nil {
		log.Println(err)
		panic("invalid filename")
	}

	return gazetteDate
}
