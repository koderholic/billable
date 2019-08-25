package config

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Data struct {
	LogPath   string `yaml:"logPath"`
}

// var config Data

func (c *Data) Init(configDir string) {

	dir, dirErr := os.Getwd()
	if dirErr != nil {
		log.Printf("Cannot set default input/output directory to the current working directory >> %s", dirErr)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath("../")
	viper.AddConfigPath(dir)
	viper.AddConfigPath(configDir)
	viper.SetDefault("inDir", dir)
	viper.SetDefault("logPath", "log.json")
	viper.WatchConfig()

	err := viper.ReadInConfig()
	if err != nil {
		// log.Write(fmt.Sprintf("\n\nfatal error: could not read from config file >>%s \n", err))
		panic(fmt.Errorf("\n fatal error: could not read from config file >>%s ", err))
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		err := viper.ReadInConfig()
		if err != nil {
			log.Printf("\n fatal error: could not read from config file >>%s ", err)
		}
		viper.Unmarshal(c)
	})

	viper.Unmarshal(c)
}
