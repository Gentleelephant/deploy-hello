package config

type Config struct {
	Pg    *PostgreSQL `yaml:"pg"`
	Redis *Redis      `yaml:"redis"`
	MySQL *MySQL      `yaml:"mysql"`
}

type NacosParams struct {
	Host      string `json:"host" yaml:"host"`
	Port      int    `json:"port" yaml:"port"`
	Namespace string `json:"namespace" yaml:"namespace"`
	DataId    string `json:"dataId" yaml:"dataId"`
	Group     string `json:"group" yaml:"group"`
	LogDir    string `json:"logDir" yaml:"logDir"`
	CacheDir  string `json:"cacheDir" yaml:"cacheDir"`
	LogLevel  string `json:"logLevel" yaml:"logLevel"`
}

type LocalNacosConfig struct {
	Nacos *NacosParams `yaml:"nacos"`
}

type PostgreSQL struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

type MySQL struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

type Redis struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
	DB   int    `json:"db" yaml:"db"`
}
