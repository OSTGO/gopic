package cmdIementaion

import (
	"fmt"
	"path"
	"strings"

	"github.com/OSTGO/gopic/conf"
	"github.com/OSTGO/gopic/utils"
)

func CmdConf() string {
	return showConfigAndPath() + "\n" + showAllPluginList() + "\n" + showActivePluginList() + "\n" + showEveryPluginHelp()
}

func showAllPluginList() string {
	return fmt.Sprintln("Supported Plugins:", utils.GetStringStringMapKey(utils.StroageHelp))
}

func showActivePluginList() string {
	return fmt.Sprintln("Actived Plugins: ", utils.GetStringUploadMapKey(utils.StroageMap))
}

func showEveryPluginHelp() string {
	strList := make([]string, 0)
	for _, v := range utils.StroageHelp {
		strList = append(strList, v)
	}
	return strings.Join(strList, "\n\n")
}

func showConfigAndPath() string {
	return "Config Path:" + fmt.Sprintln(path.Join(utils.GetHomeDir(), ".gopic.json")) + "Config Content:" + fmt.Sprintln(conf.Viper.AllSettings())
}
