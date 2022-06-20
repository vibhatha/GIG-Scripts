package helpers

import (
	"GIG-Scripts/wikipedia/wiki_web_crawler/constants"
	"bufio"
	"errors"
	"log"
	"os"
)

func LoadVisitedFromLog(visited map[string]bool) (map[string]bool, error) {
	files, err := getAllLogs(constants.VisitedLogDir)
	//if no log files exist
	if err != nil {
		return visited, err
	}
	if len(files) == 1 {
		return visited, errors.New("no log files found")
	}

	lastLog := files[len(files)-1]

	//open log file
	lastLogFile, err := os.Open(lastLog)

	if err != nil {
		return visited, err
	}
	logScanner := bufio.NewScanner(lastLogFile)
	logScanner.Split(bufio.ScanLines)

	for logScanner.Scan() {
		visited[logScanner.Text()] = true
	}
	err = lastLogFile.Close()
	if err != nil {
		return visited, err
	}
	log.Println("visited initialized from log")
	return visited, nil
}
