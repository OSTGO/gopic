package cmdIementaion

import (
	"fmt"
	"testing"
)

func TestCmdUpload(t *testing.T) {
	path := " https://pic.longtao.fun/pics/20210916/avatar.71pjc2scvak0.jpg "
	storageList := []string{"qiniu", "github"}
	args := []string{""}
	allStorage := true
	outFormat := "qiniu"
	outs := CmdUpload(storageList, args, allStorage, false, path, outFormat)
	fmt.Println("outs:", outs)
	fmt.Println("end")
}
