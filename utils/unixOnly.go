//go:build !windows
// +build !windows

package utils

import (
	"os"
)

func isLocalFile(path string) bool {
	if path == "" {
		return false
	}
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return true
	}
	return true
}
