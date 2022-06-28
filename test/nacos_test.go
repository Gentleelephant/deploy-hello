package test

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v3"
	"testing"
	"time"
)

var (
	clientConfig  constant.ClientConfig
	serverConfigs []constant.ServerConfig
	configClient  config_client.IConfigClient
)

type Config struct {
	MysqlConfig *MysqlConfig `json:"mysql" yaml:"mysql"`
	RedisConfig *RedisConfig `json:"redis" yaml:"redis"`
}

type MysqlConfig struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

type RedisConfig struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
	Db   int    `json:"db"   yaml:"db"`
}

func init() {

	clientConfig = constant.ClientConfig{
		NamespaceId:         "ff86146f-2ed9-420d-8910-3ea2d9988dff", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./log",
		CacheDir:            "./cache",
		LogLevel:            "debug",
	}

	serverConfigs = []constant.ServerConfig{
		{
			IpAddr:      "139.198.155.59",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}

	configClient, _ = clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
}

func TestGetConfig(t *testing.T) {

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "dev",
		Group:  "DEFAULT_GROUP"})
	if err != nil {
		t.Error(err)
	}

	var config = Config{}

	if err != nil {
		t.Error(err)
	}
	err = yaml.Unmarshal([]byte(content), &config)
	if err != nil {
		return
	}
	fmt.Println(config.MysqlConfig)
	fmt.Println(config.RedisConfig)
}

func listenConfig() {
	_ = configClient.ListenConfig(vo.ConfigParam{
		DataId: "dev",
		Group:  "DEFAULT_GROUP",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println(data)
		},
	})
}

func TestDynamicConfig(t *testing.T) {
	err := configClient.ListenConfig(vo.ConfigParam{
		DataId: "dev",
		Group:  "DEFAULT_GROUP",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println(data)
		},
	})
	if err != nil {
		t.Error(err)
	}
	time.Sleep(time.Second * 360)
}

func TestPushConfig(t *testing.T) {
	_, err := configClient.PublishConfig(vo.ConfigParam{
		DataId:  "dev",
		Group:   "DEFAULT_GROUP",
		Content: "",
	})
	if err != nil {
		t.Error(err)
	}
}
