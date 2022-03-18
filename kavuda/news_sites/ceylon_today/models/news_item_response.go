package models

type NewsItemsResponse struct {
	Count int                `json:"count"`
	Data  []NewsItemResponse `json:"data"`
}

type NewsItemResponse struct {
	NewsId int `json:"id"`
	Title string `json:"title"`
	Slug string `json:"slug"`
	Image string `json:"image"`
	ShortContent string `json:"short_content"`
	PublishDate string `json:"publish_date"`
	Author string `json:"author_name"`
}
