package extended_models

import (
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/models"
)

type WikipediaArticle struct {
	models.Entity
}

func (w WikipediaArticle) SetContent(content string) WikipediaArticle {
	w.Entity = w.SetAttribute("content", models.Value{
		ValueType:   ValueType.HTML,
		ValueString: content,
		Source:      "Wikipedia",
	}).AddCategory("Wikipedia")
	return w
}

func (w WikipediaArticle) GetContents() ([]models.Value, error) {
	attribute, err := w.GetAttribute("content")
	if err != nil {
		return nil, err
	}
	return attribute.GetValues(), nil
}
