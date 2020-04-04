package models

type SingleNewsItemResponse struct {
	Title       string `json:"Title"`
	AuthorName  string `json:"Author_Name"`
	PublishDate string `json:"Publish_Date"`
	HtmlContent string `json:"HTML_Content"`
	Snippet     string `json:"meta_des"`
}

type SingleNewsResponse struct {
	SuccessMessage string                   `json:"SuccessMessage"`
	ErrorMessage   string                   `json:"ErrorMessage"`
	IsSuccess      bool                     `json:"isSuccess"`
	Data           []SingleNewsItemResponse `json:"data"`
}

