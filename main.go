package main

import (
	"deploy-hello/global"
	"deploy-hello/initialization"
	"fmt"
	"time"
)

func main() {

	initialization.InitConfig()
	fmt.Printf("%+v", global.NacosConfig.Nacos)

	initialization.GetConfigFromNacos()
	fmt.Printf("%+v", global.Config.MySQL)
	time.Sleep(time.Second * 360)

}
