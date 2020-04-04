package daily_mirror

import (
	"Kavuda/models"
)

var newsSiteUrl = "http://www.dailymirror.lk/top-storys/155"
var defaultImageUrl = "http://static.dailymirror.lk/assets/uploads/advr_8570bf3960.jpg"

type DailyMirrorDecoder struct {
	models.IDecoder
}

func (d DailyMirrorDecoder) GetSourceTitle() string {
	return "Daily Mirror"
}

func (d DailyMirrorDecoder) GetDefaultImageUrl() string {
	return defaultImageUrl
}
