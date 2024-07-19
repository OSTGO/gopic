package utils

import (
	"encoding/base64"
	"golang.org/x/sys/windows"
	"net/url"
	"os"
	"syscall"
)

const (
	unknown = 0
	netPath = iota
	localPath
	base64Data
)

func isBase64(s string) bool {
	if len(s)%4 != 0 || len(s) <= 4 {
		return false
	}
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}

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

func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func CheckPath(str string) uint {

	if isBase64(str) {
		return base64Data
	} else if isLocalFile(str) {
		return localPath
	} else if isURL(str) {
		return netPath
	} else {
		return unknown
	}
}
