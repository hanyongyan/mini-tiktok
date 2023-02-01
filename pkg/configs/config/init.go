package config

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

var GlobalConfigs GlobalConfig

var ConfPath string

func init() {
	configDefaultPath := "../../pkg/configs/config/config.yaml"
	configPath := flag.String("c", configDefaultPath, "配置文件路径")
	flag.Parse()
	ConfPath = *configPath
	log.Println("use conf path: ", ConfPath)
}

func Init() {
	configFile, err := os.ReadFile(ConfPath)
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
