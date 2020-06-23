package models

type NewsSource struct {
	Link       string   `json:"link" bson:"link"`
	Categories []string `json:"categories" bson:"categories"`
}
