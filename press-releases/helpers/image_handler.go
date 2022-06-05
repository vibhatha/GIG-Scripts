package helpers

import (
	"GIG-Scripts/press-releases/constants"
	"regexp"
	"strings"
	"time"
)

var imgRE = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)

func ImageIsFound(img string) bool {
	return strings.Contains(img, "Release") && strings.Contains(img, "/images/")
}

func GetImageUrl(img string) string {
	return constants.WebsiteUrl + img
}

func GetTime(img string) (time.Time, error) {
	return time.Parse("2006.01.02", strings.Split(img, "/")[2])
}

/*
FindImages - if your img's are properly formed with doublequotes then use this, it's more efficient.
var imgRE = regexp.MustCompile(`<img[^>]+\bsrc="([^"]+)"`)
*/
func FindImages(htm string) []string {
	imgs := imgRE.FindAllStringSubmatch(htm, -1)
	out := make([]string, len(imgs))
	for i := range out {
		out[i] = imgs[i][1]
	}
	return out
}
