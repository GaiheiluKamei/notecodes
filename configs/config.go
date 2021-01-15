package configs

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Init initialize configs.
func Init() {
	viper.SetConfigFile("config.yml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config file error: %s\n", err))
	}

	if err := viper.UnmarshalExact(&Env); err != nil {
		panic(fmt.Errorf("unmarshal config file error: %s\n", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.UnmarshalExact(&Env); err != nil {
			panic(fmt.Errorf("reload config file error: %s\n", err))
		}
	})
}

// Env is global config variable.
var Env ENV

// ENV maps the config.yml
type ENV struct {
	Elastic struct {
		Addr string
		Port string
	}
	AppName                 string
	UseDifferentNameNeedTag string `mapstructure:"index"`
}
