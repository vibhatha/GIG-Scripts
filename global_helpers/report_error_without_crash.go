package global_helpers

import "log"

func ReportErrorWithoutCrash(err error) {
	if err!=nil{
		log.Println(err)
	}
}
