package utils

import (
	"os/user"
	"sync"
)

var homeDir string

var onceHomeDir sync.Once

func getHomeDir() string {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	return u.HomeDir
}

func GetHomeDir() string {
	onceHomeDir.Do(func() {
		homeDir = getHomeDir()
	})
	return homeDir
}
