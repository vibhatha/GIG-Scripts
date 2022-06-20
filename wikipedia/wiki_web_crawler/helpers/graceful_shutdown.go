package helpers

import (
	"GIG-Scripts/wikipedia/wiki_web_crawler/constants"
	"log"
	"os"
	"sync"
	"time"
)

func GracefulShutdown(queue chan string, visited map[string]bool, errorCode int) {
	log.Println("termination detected. Saving current progress...")
	currentTime := time.Now()

	var wg sync.WaitGroup
	wg.Add(2)
	go saveVisitedLog(visited, currentTime, &wg)
	go saveQueueLog(queue, visited, currentTime, errorCode, &wg)
	wg.Wait()
	os.Exit(errorCode)

}

func saveVisitedLog(visited map[string]bool, currentTime time.Time, wg *sync.WaitGroup) {
	// writing visited log
	visitedLogFilename := constants.VisitedLogDir + "visited-" + currentTime.Format(time.RFC3339) + ".log"
	visitedLog, err := os.OpenFile(visitedLogFilename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal("unable to shutdown gracefully. Error creating visited log file:", err)
	}

	for url := range visited {
		if visited[url] {
			_, err = visitedLog.WriteString(url + "\n")
		}
	}
	visitedLog.Close()
	log.Println("visited log created.")
	wg.Done()
}

func saveQueueLog(queue chan string, visited map[string]bool, currentTime time.Time, errorCode int, wg *sync.WaitGroup) {
	// writing queue log
	queueLogFilename := constants.QueueLogDir + "wiki_web_queue-" + currentTime.Format(time.RFC3339) + ".log"
	queueLog, err := os.OpenFile(queueLogFilename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal("unable to shutdown gracefully. Error creating queue log file:", err)
	}

	for {
		select {
		case url := <-queue:
			if !visited[url] {
				_, err = queueLog.WriteString(url + "\n")
			}
		default:
			queueLog.Close()
			log.Println("exiting now.")
			wg.Done()
			os.Exit(errorCode)
		}
	}
}
