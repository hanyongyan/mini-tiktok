package config

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var GlobalConfigs GlobalConfig

func Init() {
	configFile, err := os.ReadFile("../../pkg/configs/config/config.yaml")
	if err != nil {
		panic(err)
	}

	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	err = viper.ReadConfig(bytes.NewBuffer(configFile))
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	err = viper.Unmarshal(&GlobalConfigs)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
