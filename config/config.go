package config

import (
	"log"

	"github.com/Nevitash/gobsidian-master/public/configuration"
)

const (
	DEFAULT_PATH_CONFIG = "./resources/config.yaml"
)

func GetConfig() *configuration.Config {
	return configuration.GetConfig()
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
	configuration.SetConfig(config)
	return config, nil
}

func init() {
	if _, err := InitConfig(DEFAULT_PATH_CONFIG); err != nil {
		panic(err)
	}
}
