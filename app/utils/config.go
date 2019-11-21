package utils

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Http  Http  `mapstructure:"http"`
		Mysql Mysql `mapstructure:"mysql"`
	}
	Http struct {
		Port int    `mapstructure:"port"`
		Mode string `mapstructure:"mode"`
	}
	Mysql struct {
		Wdsn               string `mapstructure:"wdsn"`
		Rdsn               string `mapstructure:"rdsn"`
		MaxOpenConnections int    `mapstructure:"max_open_connections"`
		MaxIdleConnections int    `mapstructure:"max_idle_connections"`
	}
)

var (
	GlobalConfig *Config
	confPath string
)

func init(){
	flag.StringVar(&confPath, "cp", "$GOPATH/src/owner/conf",
		" set config file path")
}

func InitConfig() {
	viper.AddConfigPath(confPath)
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Unable to read config：  %s \n", err))
	}

	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		panic(fmt.Errorf("Unable to decode into struct：  %s \n", err))
	}

	fmt.Println(GlobalConfig)
}
