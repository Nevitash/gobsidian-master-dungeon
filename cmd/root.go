/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Nevitash/gobsidian-master/configuration"
	publicConfig "github.com/Nevitash/gobsidian-master/public/configuration"
	"github.com/Nevitash/gobsidian-master/public/vault"
	"github.com/spf13/cobra"
)

var (
	outputFilePath      string
	vaultPath           string
	includeFilePatterns []string
	excludeFilePatterns []string
	includePathPatterns []string
	excludePathPatterns []string
	prefixHeadings      bool
	templatePath        string
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gobsidian-master-dungeon.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&vaultPath, "vault", "v", "./vault", "Path to the Obsidian vault")
	rootCmd.Flags().StringVarP(&outputFilePath, "output", "o", "./merged-vault.md", "Output file name for the merged markdown")
	rootCmd.Flags().StringArrayVar(&includeFilePatterns, "include-file-patterns", []string{"*.md", "*.txt"}, "Glob patterns for files to include")
	rootCmd.Flags().StringArrayVar(&excludeFilePatterns, "exclude-file-patterns", []string{"*.png", "*.jpg"}, "Glob patterns for files to exclude")
	rootCmd.Flags().BoolVarP(&prefixHeadings, "prefix-headings", "p", true, "Prefix headings inside files with an extra #")
	rootCmd.Flags().StringArrayVar(&includePathPatterns, "include-path-patterns", []string{}, "Glob patterns for paths to include")
	rootCmd.Flags().StringArrayVar(&excludePathPatterns, "exclude-path-patterns", []string{"**/.*", "**/_*"}, "Glob patterns for paths to exclude")
	rootCmd.Flags().StringVarP(&templatePath, "template-path", "t", "", "Path to custom templates for rendering")
}

func mergeVault(cmd *cobra.Command, args []string) {
	var template *template.Template = nil
	var err error
	if templatePath != "" {
		template, err = publicConfig.LoadTemplateFromFile(templatePath)
		if err != nil {
			fmt.Printf("Error loading template from %s: %v\n", templatePath, err)
			return
		}
	}
	config, err := vault.CreateConfig(vaultPath,
		includeFilePatterns,
		excludeFilePatterns,
		includePathPatterns,
		excludePathPatterns,
		&configuration.Flags{PrefixHeadings: prefixHeadings},
		template)
	if err != nil {
		fmt.Printf("Error creating vault config: %v\n", err)
		return
	}
	fmt.Printf("Config %+v\n", config)
	vault.MergeToFileWithConfig(outputFilePath, config)
	outputFilePath = filepath.Clean(outputFilePath)
	outputFilePath, err = filepath.Abs(outputFilePath)
	if err != nil {
		fmt.Printf("Error getting absolute path for output file: %v\n", err)
		return
	}
	fmt.Printf("Merged vault written to %s\n", outputFilePath)
}
