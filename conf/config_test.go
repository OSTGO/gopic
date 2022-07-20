package conf

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	qiniu := Viper.GetStringMap("qiniu")
	//fmt.Println(viper.AllKeys())
	fmt.Println(qiniu)
	if qiniu["active"] == false {
		fmt.Println("fuck")
	}
	fmt.Println(qiniu)
	fmt.Println(qiniu["bucket"])
}
