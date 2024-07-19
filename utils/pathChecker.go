package utils

import (
	"encoding/base64"
	"net/url"
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
