package initialization

import (
	"deploy-hello/global"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v3"
)

func GetConfigFromNacos() {
	clientConfig := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Nacos.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              global.NacosConfig.Nacos.LogDir,
		CacheDir:            global.NacosConfig.Nacos.CacheDir,
		LogLevel:            global.NacosConfig.Nacos.LogLevel,
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      global.NacosConfig.Nacos.Host,
			ContextPath: "/nacos",
			Port:        uint64(global.NacosConfig.Nacos.Port),
			Scheme:      "http",
		},
	}

	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.Nacos.DataId,
		Group:  global.NacosConfig.Nacos.Group,
	})
	if err != nil {
		panic(err)
	}
	if err = parseConfig(content, global.Config); err != nil {
		panic(err)
	}

	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: global.NacosConfig.Nacos.DataId,
		Group:  global.NacosConfig.Nacos.Group,
		OnChange: func(namespace, group, dataId, data string) {
			// 配置修改，重新加载配置
			fmt.Println("====================监听到配置修改====================")
			err := parseConfig(data, global.Config)
			if err != nil {
				return
			}
		},
	})

}

func parseConfig(content string, config interface{}) error {
	fmt.Println(content)
	// 重新初始化DB
	err := yaml.Unmarshal([]byte(content), config)
	if err != nil {
		return err
	}
	// 初始化DB
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//global.DbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	return err
	//}
	return err
}
