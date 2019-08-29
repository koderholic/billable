package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Data struct {
	LogPath string `yaml:"logPath"`
	Port string `yaml:"port"`
}

// var config Data

func (c *Data) Init(configDir string) error {

	dir, dirErr := os.Getwd()
	if dirErr != nil {
		log.Printf("Cannot set default input/output directory to the current working directory >> %s", dirErr)
		return dirErr
	}

	if configDir != "" {
		viper.AddConfigPath(configDir)
	} else {
		viper.AddConfigPath("../")
		viper.AddConfigPath(dir)
	}

	viper.SetConfigName("config")
	viper.SetDefault("logPath", "log.json")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("\n fatal error: could not read from config file >> %s ", err)
		return err
	}

	viper.Unmarshal(c)
	return nil
}
