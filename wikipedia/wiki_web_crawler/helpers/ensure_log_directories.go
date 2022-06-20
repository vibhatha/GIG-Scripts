package helpers

import (
	"GIG-Scripts/wikipedia/wiki_web_crawler/constants"
	"github.com/lsflk/gig-sdk/libraries"
)

func EnsureLogDirectories() error {
	err := libraries.EnsureDirectory(constants.LogDir)
	if err != nil {
		return err
	}
	err = libraries.EnsureDirectory(constants.QueueLogDir)
	if err != nil {
		return err
	}
	err = libraries.EnsureDirectory(constants.VisitedLogDir)
	if err != nil {
		return err
	}
	return nil
}
