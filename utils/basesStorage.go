package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	netPath = iota
	localPath
	base64Data
)

type BaseStorage struct {
	ImageList []*Image
	token     string
}

type Image struct {
	Path       string
	Title      string
	ImageType  string
	PathType   uint
	OutBytes   []byte
	OutBase64  string
	InName     string
	OutSuffix  string
	FolderName string
	OutName    string
	md5        [md5.Size]byte
}

var bMap map[string]*BaseStorage

func NewBaseStorage(paths []string, nameReserve bool) *BaseStorage {
	token, err := json.Marshal(paths)
	if err != nil {
		panic(err)
	}
	var b *BaseStorage
	var ok bool
	if b, ok = bMap[string(token)]; !ok {
		b = &BaseStorage{token: string(token)}
	}
	err = b.Generate(paths, nameReserve)
	if err != nil {
		panic(err)
	}
	return b
}

func (b *BaseStorage) Destory() {
	delete(bMap, b.token)
}

func (b *Image) setPath(rawImagePath string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()
	// 处理标题，找到第一个空格split，[0]是imagePath、[1]strip("\"")是标题
	rawImagePathList := strings.Split(rawImagePath, " ")
	imagePath := rawImagePathList[0]
	var title string
	if len(rawImagePathList) >= 2 {
		title = fmt.Sprintf(strings.Join(rawImagePathList[1:], " "))
	}
	if len(imagePath) > 5 && imagePath[:5] == "data:" {
		pathRaw := strings.Split(imagePath, "base64,")
		if len(pathRaw) < 1 {
			panic(errors.New("unknown data type"))
		}
		b.PathType = base64Data
		b.InName = "base64"
		b.Path = pathRaw[1]
		return nil
	}
	b.Title = title
	b.Path = imagePath
	_, picName := path.Split(imagePath)
	picNameSplit := strings.Split(picName, "_gopic_")
	if len(picNameSplit) >= 2 {
		b.InName = strings.Split(picNameSplit[len(picNameSplit)-1], ".")[0]
	} else {
		b.InName = strings.Split(picName, ".")[0]
	}
	_, err = os.Stat(imagePath)
	if os.IsNotExist(err) {
		b.PathType = netPath
	} else {
		b.PathType = localPath
	}
	return nil
}

func (b *Image) processData() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()
	var byteArrayData []byte
	switch b.PathType {
	case netPath:
		byteArrayData, err = netPictureData(b.Path)
	case localPath:
		byteArrayData, err = localPictureData(b.Path)
	case base64Data:
		byteArrayData, err = base64PictureData(b.Path)
	default:
		panic("Image type error!")
	}
	b.OutBytes = byteArrayData
	b.ImageType = CheckImageType(byteArrayData)
	b.OutBase64 = base64.StdEncoding.EncodeToString(byteArrayData)
	return nil
}

func (b *Image) setOthers(nameReserve bool) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()
	b.FolderName = getCurrentYear()
	b.md5 = md5.Sum(b.OutBytes)
	if nameReserve {
		b.OutName = bytesRaw2String(b.md5[:]) + "_gopic_" + b.InName + "." + b.ImageType
	} else {
		b.OutName = bytesRaw2String(b.md5[:]) + "_gopic_" + "." + b.ImageType
	}
	b.OutSuffix = path.Join(b.FolderName, b.OutName)
	return nil
}

func (b *BaseStorage) Generate(paths []string, nameReserve bool) error {
	flag := 0
	var wg sync.WaitGroup
	for k, pa := range paths {
		wg.Add(1)
		go func(index int, path string) {
			i := &Image{}
			i.setPath(path)
			i.processData()
			i.setOthers(nameReserve)
			if index == flag {
				b.ImageList = append(b.ImageList, i)
			}
			flag++
			wg.Done()
		}(k, pa)
		wg.Wait()
	}
	return nil
}

//2019->19;2020->20
func getCurrentYear() string {
	return strconv.Itoa(time.Now().Year() - 2000)
}

func bytesRaw2String(bs []byte) string {
	var out string
	for _, b := range bs {
		out += fmt.Sprint(b)
	}
	return out
}

func netPictureData(netPath string) ([]byte, error) {
	res, err := http.Get(netPath)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	originData, _ := ioutil.ReadAll(res.Body)
	return originData, nil
}

func localPictureData(localPath string) ([]byte, error) {
	originData, err := ioutil.ReadFile(localPath)
	if err != nil {
		return nil, err
	}
	return originData, nil
}

func base64PictureData(localPath string) ([]byte, error) {
	originData, err := base64.StdEncoding.DecodeString(localPath)
	return originData, err
}
