package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	InitConfig = inticonfig
)

func inticonfig(filename string, additionalDirs []string) (err error) {
	viper.SetConfigName(filename)
	viper.AddConfigPath(".")

	for _, dir := range additionalDirs {
		viper.AddConfigPath(dir)
	}

	err = viper.ReadInConfig()
	if err != nil {
		log.Println("Failed to read config file")
	}

	viper.ConfigFileUsed()
	viper.WatchConfig()
	return
}