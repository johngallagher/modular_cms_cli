/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
	"github.com/joyfulprogramming/modular_cms_cli/modular"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a page",
	Long:  "Edit a page",
	Run: func(cmd *cobra.Command, args []string) {
		filePath, _ := cmd.Flags().GetString("file")

		fmt.Print("\033[H\033[2J") // Clear screen
		p := tea.NewProgram(modular.InitialModel(filePath))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error running program: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	pageCmd.AddCommand(editCmd)
	editCmd.Flags().StringP("file", "f", "src/index.md", "Path to the markdown file")
}
