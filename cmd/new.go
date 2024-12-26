/*
Copyright © 2024 Joyful Programming
*/
package cmd

import (
	"github.com/spf13/cobra"
	"modular_cms_cli/modular/commands"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new SITE_PATH",
	Short: "Create a new Modular site",
	Long:  "Create a new Modular site at the specified path",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		commands.New(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
