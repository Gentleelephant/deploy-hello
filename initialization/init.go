package initialization

import (
	"deploy-hello/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// InitConfig 初始化配置
func InitConfig() {

	logger := global.Logger

	configFile := global.ConfigFile
	v := viper.New()
	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		logger.Error("viper read config err", zap.Error(err))
	}
	if err := v.Unmarshal(global.NacosConfig); err != nil {
		logger.Error("viper unmarshal config err", zap.Error(err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		logger.Info("config file changed...")
		err := v.Unmarshal(global.NacosConfig)
		if err != nil {
			logger.Error("unmarshal emote nacos config err", zap.Error(err))
		}
	})
}

// 初始化配置
func InitVariable() {

}
