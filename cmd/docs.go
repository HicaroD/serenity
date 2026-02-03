package cmd

import (
	"github.com/serenitysz/serenity/internal/cmds/docs"
	"github.com/spf13/cobra"
)

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Open Serenity documentation in the browser quickly",
	RunE: func(cmd *cobra.Command, args []string) error {
		return docs.Open()
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
