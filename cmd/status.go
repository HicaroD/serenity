package cmd

import (
	"github.com/serenitysz/serenity/internal/cmds/status"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display the current status of Serenity",
	RunE: func(cmd *cobra.Command, args []string) error {
		return status.Get()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
