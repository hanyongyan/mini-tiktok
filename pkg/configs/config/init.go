package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
)

func Init() {
	//	path, _ := os.Getwd()
	path, _ := filepath.Abs("./../../")
	log.Println("workspace path", path)

	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(fmt.Sprintf("%v/pkg/configs/config", path))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
