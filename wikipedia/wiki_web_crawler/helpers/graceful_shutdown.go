package helpers

import (
	"log"
	"os"
	"time"
)

const QueueLog = "tmp/wiki_web_queue-"

func GracefulShutdown(queue chan string, errorCode int) {
	log.Println("termination detected. Saving current progress...")
	currentTime := time.Now()
	filename := QueueLog + currentTime.Format(time.RFC3339) + ".log"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal("unable to shutdown gracefully. Error creating log file:", err)
	}

	for {
		select {
		case url := <-queue:
			_, err = file.WriteString(url + "\n")
		default:
			file.Close()
			log.Println("exiting now.")
			os.Exit(errorCode)
		}
	}
}
