package cmdIementaion

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/OSTGO/gopic/utils"
)

func CmdConvert(covertPath, outDir, outFormat string, allStorage, nameReserve, recurse bool, storageList []string) error {
	if covertPath == "" {
		covertPath = "./"
	}
	if outDir == "" {
		return errors.New("not set outDir")
	}
	if allStorage {
		storageList = utils.GetStringUploadMapKey(utils.StroageMap)
	}
	if storageList == nil || len(storageList) == 0 {
		return errors.New("not chose storage")
	}
	if outFormat == "" {
		outFormat = storageList[0]
	}
	err := convert(covertPath, outDir, outFormat, storageList, nameReserve)
	return err
}

//对单个文件转换
func convertFile(filePath, outPath, outFormat string, stotageList []string, nameReserve bool) error {
	fmt.Println("sourcePath:", filePath, " -> ", "targetPath:", outPath)
	fd, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer fd.Close()
	rawByteData, err := ioutil.ReadAll(fd)
	if err != nil {
		panic(err)
	}
	outByteData, err := convertByteData(rawByteData, stotageList, outFormat, nameReserve)
	if err != nil {
		return err
	}
	_, err = os.Stat(filepath.Dir(outPath))
	if err != nil && !os.IsExist(err) {
		err = os.MkdirAll(filepath.Dir(outPath), os.ModePerm)
	}
	err = ioutil.WriteFile(outPath, outByteData, 0644)
	return err
}

//对字节数组转换
func convertByteData(rawByteData []byte, stotageList []string, outFormat string, nameReserve bool) ([]byte, error) {
	outStringData, err := convertStringData(string(rawByteData), stotageList, outFormat, nameReserve)
	return []byte(outStringData), err
}

func convertStringData(rawData string, stotageList []string, outFormat string, nameReserve bool) (string, error) {
	picList, err := findPicList(rawData)
	if err != nil {
		return "", err
	}
	picList, err = uploadPicList(picList, stotageList, outFormat, nameReserve)
	out, err := replaceData(rawData, picList)
	return out, err
}

//获取图片列表
func findPicList(rawData string) ([][]string, error) {
	reg1 := regexp.MustCompile(`!\[+.+]\(+.+\)`)
	if reg1 == nil {
		return nil, errors.New("regexp err")
	}
	result1 := reg1.FindAllStringSubmatch(rawData, -1)
	return result1, nil
}

//上传文件列表
func uploadPicList(picList [][]string, stotageList []string, outFormat string, nameReserve bool) ([][]string, error) {
	picList1D := make([]string, 0, len(picList))
	for i, v := range picList {
		kk := append(strings.Split(v[0], "("))
		picList[i] = append(v, kk[0])
		kk1List := strings.Split(strings.TrimSuffix(kk[1], ")"), " ")
		picPath := kk1List[0]
		var title string
		if len(kk1List) >= 2 {
			title = fmt.Sprintf(strings.Join(kk1List[1:], " "))
		}
		picList1D = append(picList1D, picPath)
		picList[i] = append(picList[i], title)
	}
	// outList1D := realUploadPicList(picList1D)
	outMap, errMap := NewUpload(stotageList, picList1D, nameReserve)
	if len(errMap) != 0 {
		for k, v := range errMap {
			fmt.Printf("%v:%v\n", k, v)
		}
		return nil, errors.New("uploadPicList error")
	}
	for i, v := range outMap[outFormat] {
		picList[i][1] = picList[i][1] + "(" + v + " " + picList[i][2] + ")"
	}
	return picList, nil
}

func replaceData(rawData string, picList [][]string) (string, error) {
	for _, v := range picList {
		rawData = strings.ReplaceAll(rawData, v[0], v[1])
	}
	return rawData, nil
}

func convert(inPath, outPath, outFormat string, stotageList []string, nameReserve bool) error {
	s, err := os.Stat(inPath)
	if err != nil {
		return err
	}
	if s.IsDir() {
		err = convertDir(inPath, outPath, outFormat, stotageList, nameReserve)

	} else {
		err = convertFile(inPath, outPath, outFormat, stotageList, nameReserve)
	}
	return err
}

func convertDir(dirPath, outPath, outFormat string, stotageList []string, nameReserve bool) error {
	//dirPath2outPath := filepath.Rel()
	err := filepath.Walk(dirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			relativePath, _ := filepath.Rel(dirPath, path)
			outP := filepath.Join(outPath, relativePath)
			if info.IsDir() {
				_, err = os.Stat(outP)
				if err != nil && !os.IsExist(err) {
					err = os.MkdirAll(filepath.Dir(outPath), os.ModePerm)
				}
				return nil
			}
			err = convertFile(path, outP, outFormat, stotageList, nameReserve)
			if err != nil {
				return err
			}
			if info.Size() <= 5 {
				return nil
			}
			return nil
		})
	return err
}
