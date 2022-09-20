package cmdIementaion

import (
	"io/ioutil"
	"path"

	"github.com/OSTGO/gopic/staic"
	"github.com/OSTGO/gopic/utils"
)

func CmdInit() (string, string) {
	data := staic.Config
	configPath := path.Join(utils.GetHomeDir(), ".gopic.json")
	err := ioutil.WriteFile(configPath, []byte(data), 0777)
	if err != nil {
		panic(err)
	}
	return data, configPath
}
