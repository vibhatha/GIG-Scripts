package decoders

import "github.com/lsflk/gig-sdk/models"

type MyLocalDecoder interface {
	DecodeToEntity(record []string, source string) models.Entity
}
