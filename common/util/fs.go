package util

import (
	"os"
)

// CheckExist check if a file or dir is Exist.
func CheckExist(filepath string) bool {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false
	}
	return true
}

// IsDir check if the path is a file, false when not exist or is a dir.
func IsDir(path string) bool {
	if CheckExist(path) {
		stat, _ := os.Stat(path)
		return stat.IsDir()
	}
	return false
}

// IsFile check if a path a file, false when not exist or is a file.
func IsFile(path string) bool {
	if CheckExist(path) {
		stat, _ := os.Stat(path)
		return !stat.IsDir()
	}
	return false
}
