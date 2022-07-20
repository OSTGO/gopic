package cmdIementaion

import (
	"fmt"
	"testing"
)

func TestCmdUpload(t *testing.T) {
	pathList := []string{"../1.png"}
	storageList := []string{"qiniu", "github"}
	args := []string{"../2.png"}
	allStorage := false
	outFormat := "qiniu"
	outs := CmdUpload(pathList, storageList, args, allStorage, outFormat)
	fmt.Println(outs)
	fmt.Println("end")
}
