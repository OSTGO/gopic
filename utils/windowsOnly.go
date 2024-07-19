//go:build windows
// +build windows

package utils

import (
	"golang.org/x/sys/windows"
	"os"
	"syscall"
)

func isLocalFile(path string) bool {
	if path == "" {
		return false
	}
	_, err := os.Stat(path)
	if err != nil {
		if pathErr, ok := err.(*os.PathError); ok {
			if errno, ok := pathErr.Err.(syscall.Errno); ok && errno == windows.ERROR_INVALID_NAME {
				return false
			}
		}
		return !os.IsNotExist(err)
	}
	return true
}
