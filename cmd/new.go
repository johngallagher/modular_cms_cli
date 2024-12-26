/*
Copyright © 2024 Joyful Programming
*/
package cmd

import (
	"modular_cms_cli/modular/commands"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new SITE_PATH",
	Short: "Create a new Modular site",
	Long:  "Create a new Modular site at the specified path",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		commands.CreateNewSite(args[0])
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
