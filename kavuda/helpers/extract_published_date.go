package helpers

import (
	"log"
	"time"
)

/*
return date from string or if the date layout is different return current date
 */
func ExtractPublishedDate(layout string, timeString string) time.Time {
	loc, _ := time.LoadLocation("Asia/Colombo")
	t, err := time.ParseInLocation(layout, timeString,loc)
	if err != nil {
		log.Println("error in date", err)
		t = time.Now().In(loc)
	}
	return t
}
