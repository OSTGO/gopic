package utils

import (
	"fmt"
	"testing"
)

func TestNewBaseStorage(t *testing.T) {
	n := NewBaseStorage()
	//n.Generate("1.png")
	fmt.Println(n)

}
