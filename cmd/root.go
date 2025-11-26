/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/Nevitash/gobsidian-master-dungeon/config"
	"github.com/Nevitash/gobsidian-master/public/file"
	"github.com/spf13/cobra"
)

var (
	outputFilePath string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gobsidian-master-dungeon -v [/path/to/vault] -o [output-file-path]",
	Short: "Merges all files in a folder into a single markdown file for Obsidian",
	Run:   mergeVault,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	config := config.GetConfig()
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gobsidian-master-dungeon.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&config.VaultPath, "vault", "v", "./", "Path to the Obsidian vault")
	rootCmd.Flags().StringVarP(&outputFilePath, "output", "o", "./merged-vault.md", "Output file name for the merged markdown")
	rootCmd.Flags().StringArrayVar(&config.IncludeFilePatterns, "include-file-patters", []string{"*.md", "*.txt"}, "Glob patterns for files to include")
	rootCmd.Flags().StringArrayVar(&config.ExcludeFilePatterns, "exclude-file-patterns", []string{"*.png", "*.jpg"}, "Glob patterns for files to exclude")
	rootCmd.Flags().BoolVarP(&config.Flags.PrefixHeadings, "prefix-headings", "p", true, "Prefix headings inside files with an extra #")
	rootCmd.Flags().StringArrayVar(&config.IncludePathPatterns, "include-path-patterns", []string{}, "Glob patterns for paths to include")
	rootCmd.Flags().StringArrayVar(&config.ExcludePathPatterns, "exclude-path-patterns", []string{"**/_*", "**/.*"}, "Glob patterns for paths to exclude")
}

func mergeVault(cmd *cobra.Command, args []string) {
	config := config.GetConfig()
	vault, err := file.LoadVaultFile(outputFilePath, config)
	if err != nil {
		panic(err)
	}
	file.CombineVault(vault, config)
}
