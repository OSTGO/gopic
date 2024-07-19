package cmdIementaion

import (
	"fmt"
	"sync"

	"github.com/OSTGO/gopic/utils"

	_ "github.com/OSTGO/gopic/plugin"
)

func CmdUpload(storageList, args []string, allStorage, nameReserve bool, path, outFormat string) string {
	if path == "" && args == nil {
		return "path is nil"
	}
	pathList := make([]string, 0, len(args)+1)
	if args != nil {
		pathList = append(args, path)
	} else {
		pathList = []string{path}
	}
	if allStorage || storageList == nil {
		storageList = utils.GetStringUploadMapKey(utils.StroageMap)
	}
	outMap, errMap := NewUpload(storageList, utils.StrimList(pathList), nameReserve)
	if len(errMap) != 0 {
		for k, v := range errMap {
			fmt.Printf("%v:%v\n", k, v)
		}
		return ""
	}
	if outFormat == "" {
		outFormat = storageList[0]
	}
	urlList := ""
	for _, v := range outMap[outFormat] {
		urlList = urlList + v + "\n"
	}
	return urlList
}

var errLock sync.Mutex

func NewUpload(stroages []string, paths []string, nameReserve bool) (map[string][]string, map[string][]error) {
	errMapList := make(map[string][]error, 0)
	outMapList := make(map[string][]string, len(stroages))
	bb := utils.NewBaseStorage(paths, nameReserve)
	defer bb.Destory()
	var wg sync.WaitGroup
	for _, stroage := range stroages {
		wg.Add(1)
		go func(_stroage string, paths []string) {
			out, err := stroageUpload(_stroage, paths, nameReserve)
			if len(err) != 0 {
				errLock.Lock()
				errMapList[_stroage] = err
				errLock.Unlock()
			} else {
				errLock.Lock()
				outMapList[_stroage] = out
				errLock.Unlock()
			}
			wg.Done()
		}(stroage, paths)
	}
	wg.Wait()
	return outMapList, errMapList
}

// need performance optimization
func stroageUpload(stroage string, paths []string, nameReserve bool) ([]string, []error) {
	st := utils.StroageMap[stroage]
	base := utils.NewBaseStorage(paths, nameReserve)
	st.SetPicList(paths, nameReserve)
	var wg sync.WaitGroup
	flag := 0
	//flag := make([]chan bool, len(base.ImageList), len(base.ImageList))
	//flag[0] <- true
	outList := make([]string, 0, len(base.ImageList))
	errList := make([]error, 0, len(base.ImageList))
	for k, v := range base.ImageList {
		wg.Add(1)
		go func(index int, im *utils.Image) {
			outURL, err := st.Upload(im)
			//select {
			//case <-flag[index]:
			//	outList = append(outList, outURL)
			//	if err != nil {
			//		errList = append(errList, err)
			//	}
			//	if index+1 < len(base.ImageList) {
			//		flag[index+1] <- true
			//	}
			//}
			for flag != index {
			}
			outList = append(outList, outURL)
			if err != nil {
				errList = append(errList, err)
			}
			flag++
			wg.Done()
		}(k, v)
	}
	wg.Wait()
	return outList, errList
}
