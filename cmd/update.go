package cmd

import (
	"github.com/serenitysz/serenity/internal/version"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Serenity to the latest version",
	RunE: func(cmd *cobra.Command, args []string) error {
		return version.Update()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
