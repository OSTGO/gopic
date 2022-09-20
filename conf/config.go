package conf

import (
	"fmt"
	"path"

	"github.com/OSTGO/gopic/utils"
	"github.com/spf13/viper"
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
