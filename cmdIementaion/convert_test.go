package cmdIementaion

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func Test_findPicList(t *testing.T) {
	data, err := ioutil.ReadFile("tem.md")
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(data))
	m, err := findPicList(string(data))
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
}

func Test_uploadPicList(t *testing.T) {
	picList := make([][]string, 1, 1)
	picList[0] = []string{"![image-20220802162132491](https://pic.longtao.fun/pics/22/21692137714958140223586238239226443276_image-20220802162132491.png)"}
	stotageList := []string{"samba"}
	outFormat := "samba"
	out, err := uploadPicList(picList, stotageList, outFormat)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

func Test_convert(t *testing.T) {
	err := convertFile("tem.md", "tem.md", "samba", []string{"samba"})
	if err != nil {
		panic(err)
	}
}

func Test_convertDir(t *testing.T) {
	//inPath, _ := filepath.Abs("/home/longtao/workspace/document/blog")
	//outPath, _ := filepath.Abs("/home/longtao/temp/blog")
	inPath, _ := filepath.Abs("./tem.md")
	outPath, _ := filepath.Abs("../tem2/tem.md")
	err := convert(inPath, outPath, "samba", []string{"samba"})
	if err != nil {
		panic(err)
	}
	//err := convertDir(inPath, outPath, "samba", []string{"samba"})
	//if err != nil {
	//	panic(err)
	//}
}
