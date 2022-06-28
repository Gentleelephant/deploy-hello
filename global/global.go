package global

import (
	"deploy-hello/config"
	"gorm.io/gorm"
)

var (
	NacosConfig = &config.LocalNacosConfig{}
	Config      = &config.Config{}
	DbConn      = &gorm.DB{}
	ConfigFile  = "config/config.yaml"
)
