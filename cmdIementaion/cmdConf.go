package cmdIementaion

import (
	"fmt"
	"gopic/conf"
	"gopic/utils"
	"path"
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
	return fmt.Sprintln(utils.StroageHelp)
}

func showConfigAndPath() string {
	return "Config Path:" + fmt.Sprintln(path.Join(utils.GetHomeDir(), ".gopic.json")) + "Config Content:" + fmt.Sprintln(conf.Viper.AllSettings())
}
