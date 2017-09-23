package utils

import "os"

func IsDirExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}

	return false
}

func ArrayContainer(array []string, found string) bool {
	for _, v := range array {
		if v == found {
			return true
		}
	}

	return false
}
