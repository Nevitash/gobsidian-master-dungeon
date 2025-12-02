/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/Nevitash/gobsidian-master-dungeon/cmd"
	"github.com/Nevitash/gobsidian-master/public/vault"
)

func main() {
	vault.MergeToFile("./vault", "./merged_vault.md", nil, nil, nil, nil, nil, nil)
	cmd.Execute()
}
