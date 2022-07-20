package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"gopic/utils"
	"path"
)

var Viper viper.Viper

func init() {
	Viper = *viper.New()
	configPath := path.Join(utils.GetHomeDir(), ".gopic.json")
	Viper.SetConfigFile(configPath)
	if err := Viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}
