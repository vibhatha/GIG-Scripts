package extended_models

import (
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/models"
	"time"
)

type Tender struct {
	models.Entity
}

func (t *Tender) SetTenderTitle(title string, source string, sourceDate time.Time) *Tender {
	t.SetTitle(models.Value{
		ValueType:   ValueType.String,
		ValueString: title,
		Source:      source,
		Date:        sourceDate,
	})
	return t
}
func (t *Tender) SetCompany(companyName string) *Tender {
	t.SetAttribute("Company", models.Value{
		ValueType:   ValueType.String,
		ValueString: companyName,
	})
	return t
}

func (t Tender) GetCompany() string {
	attribute, err := t.GetAttribute("Company")
	if err != nil {
		return ""
	}
	return attribute.GetValue().GetValueString()
}

func (t *Tender) SetClosingDate(closingDate time.Time) *Tender {
	t.SetAttribute("Closing Date", models.Value{
		ValueType:   ValueType.Date,
		ValueString: closingDate.String(),
	})
	return t
}

func (t *Tender) SetLocation(location string) *Tender {
	t.SetAttribute("Location", models.Value{
		ValueType:   ValueType.String,
		ValueString: location,
	})
	return t
}

func (t Tender) GetLocation() string {
	attribute, err := t.GetAttribute("Location")
	if err != nil {
		return ""
	}
	return attribute.GetValue().GetValueString()
}

func (t *Tender) SetDescription(description string) *Tender {
	t.SetAttribute("Description", models.Value{
		ValueType:   ValueType.String,
		ValueString: description,
	})
	return t
}

func (t *Tender) SetTenderValue(value string) *Tender {
	t.SetAttribute("Description", models.Value{
		ValueType:   ValueType.Number,
		ValueString: value,
	})
	return t
}
