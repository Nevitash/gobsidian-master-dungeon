/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/Nevitash/gobsidian-master-dungeon/cmd"
	"github.com/Nevitash/gobsidian-master-dungeon/config"
)

const (
	DEFAULT_PATH_CONFIG = "./resources/config.yaml"
)

func main() {
	if _, err := config.InitConfig(DEFAULT_PATH_CONFIG); err != nil {
		panic(err)
	}
	cmd.Execute()
}
