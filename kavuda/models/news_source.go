package models

import "time"

type NewsItem struct {
	Date     time.Time `json:"date" bson:"date"`
	Title    string    `json:"title" bson:"title"`
	Snippet  string    `json:"snippet" bson:"snippet"`
	Link     string    `json:"link" bson:"link"`
	Content  string    `json:"content" bson:"content"`
	Author   string    `json:"author"`
	ImageURL string    `json:"image_url" bson:"image_url"`
}
