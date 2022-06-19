package helpers

import (
	"flag"
	"log"
	"os"
)

func CheckArgs() []string {
	flag.Parse()
	args := flag.Args()
	log.Println(args)
	if len(args) < 1 {
		log.Println("starting url not specified")
		os.Exit(1)
	}
	return args
}
