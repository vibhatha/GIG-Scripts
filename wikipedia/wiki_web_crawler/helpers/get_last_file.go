package helpers

import "errors"

func getLastFile(files []string) (string, error) {
	if len(files) == 1 {
		return "", errors.New("no log files found")
	}
	return files[len(files)-1], nil
}
