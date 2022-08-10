package cmdIementaion

import (
	"fmt"
	"testing"
)

func TestCmdUpload(t *testing.T) {
	path := "../1.png"
	storageList := []string{"qiniu", "github"}
	args := []string{"../2.png"}
	allStorage := false
	outFormat := "qiniu"
	outs := CmdUpload(storageList, args, allStorage, false, path, outFormat)
	fmt.Println(outs)
	fmt.Println("end")
}
