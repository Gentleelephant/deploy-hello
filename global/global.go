package global

import (
	"deploy-hello/config"
	"deploy-hello/global/zap"
	"gorm.io/gorm"
)

var (
	NacosConfig = &config.LocalNacosConfig{}
	Config      = &config.Config{}
	DbConn      = &gorm.DB{}
	ConfigFile  = "config/config.yaml"
	Logger      = zap.GetLogger()
)
