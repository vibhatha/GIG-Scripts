package ada_derana

import (
	"GIG-Scripts/kavuda/models"
)

var newsSiteUrl = "http://www.adaderana.lk/hot-news/"
var defaultImageUrl = "http://www.adaderana.lk/newstyle2017/assets/images/adaderana-header-logo-2020.png"

type AdaDeranaDecoder struct {
	models.IDecoder
}

func (d AdaDeranaDecoder) GetSourceTitle() string {
	return "Ada Derana"
}

func (d AdaDeranaDecoder) GetDefaultImageUrl() string {
	return defaultImageUrl
}
