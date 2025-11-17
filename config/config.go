package config

import (
	"log"

	"github.com/Nevitash/gobsidian-master/public/configuration"
)

var config *configuration.Config

func GetConfig() *configuration.Config {
	return config
}

func InitConfig(configPath string) (*configuration.Config, error) {
	config, err := configuration.LoadConfig(configPath)
	if err != nil {
		log.Printf("No config file found. Creating a new one: %v", configPath)
		config = configuration.NewDefaultConfig(configPath)
		err = configuration.SaveConfig(config)
	}
	if err != nil {
		log.Fatalf("Failed to create config file: %v", err)
		return nil, err
	}
	log.Printf("Config loaded: %+v", config)
	return config, nil
}
