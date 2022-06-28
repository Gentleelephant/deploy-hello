package initialization

import (
	"deploy-hello/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// InitConfig 初始化配置
func InitConfig() {
	configFile := "config/config.yaml"
	v := viper.New()
	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(global.NacosConfig); err != nil {
		panic(err)
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("====================配置修改====================")
		err := v.Unmarshal(global.NacosConfig)
		if err != nil {
			panic(err)
		}
	})
}

// 初始化配置
func InitVariable() {

}
