/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "modular_cli",
	Short: "CLI tool to build static sites with components in seconds",
	Long:  "Convention over configuration for static sites. GUI landing page builders are clumsy, slow, bug ridden and annoying to use for engineers. Static site builders are better, but setting one up requires hours of research and fiddling with hundreds of options. Modular is a CLI driven CMS that allows you to build fast, beautiful landing pages with your favorite component system in seconds.",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	pageCmd.Run(cmd, args)
	// },
}

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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.modular_cms_cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
