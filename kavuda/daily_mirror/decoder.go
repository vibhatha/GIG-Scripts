package daily_mirror

import (
	"GIG-Scripts/kavuda/models"
)

var defaultImageUrl = "http://static.dailymirror.lk/assets/uploads/advr_8570bf3960.jpg"
var newsSources = []models.NewsSource{
	{Link: "http://www.dailymirror.lk/top-storys/155", Categories: []string{"Local"}},
	{Link: "http://www.dailymirror.lk/latest-news/108", Categories: []string{"Local"}},
	{Link: "http://www.dailymirror.lk/news/209", Categories: []string{"Local"}},
}

type DailyMirrorDecoder struct {
	models.IDecoder
}

func (d DailyMirrorDecoder) GetSourceTitle() string {
	return "Daily Mirror"
}

func (d DailyMirrorDecoder) GetDefaultImageUrl() string {
	return defaultImageUrl
}
