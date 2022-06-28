package server

import (
	"deploy-hello/initialization"
	"time"
)

func Start() {

	initialization.InitConfig()

	initialization.GetConfigFromNacos()

	time.Sleep(time.Second * 30)

}
