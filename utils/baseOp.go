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

func StrimList(l []string) []string {
	if l == nil || len(l) == 0 {
		return l
	}
	pointer := 0
	for _, v := range l {
		if v != "" {
			l[pointer] = v
			pointer++
		}
	}
	l = l[:pointer]
	return l
}

func DeleteAfterLastCharacter(s string, character string) string {
	if s == "" || len(s) == 0 || len(character) == 0 {
		return s
	}
	c := character[0]
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			s = s[:i]
			return s
		}
	}
	return s
}
