package helpers

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func LoadQueueFromLog(queue chan string) error {
	files, err := getAllLogs()
	//if no log files exist
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return errors.New("no log files found")
	}

	lastLog := files[len(files)-1]

	//open log file
	lastLogFile, err := os.Open(lastLog)

	if err != nil {
		return err
	}
	logScanner := bufio.NewScanner(lastLogFile)
	logScanner.Split(bufio.ScanLines)

	for logScanner.Scan() {
		go func(url string) { queue <- url }(logScanner.Text())
	}
	err = lastLogFile.Close()
	if err != nil {
		return err
	}
	log.Println("queue initialized from log")
	return nil
}

func getAllLogs() ([]string, error) {
	var files []string
	err := filepath.Walk("tmp/", func(path string, info os.FileInfo, err error) error {
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
