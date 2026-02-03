package cmd

import (
	"os"

	"github.com/serenitysz/serenity/internal/exception"
	"github.com/serenitysz/serenity/internal/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: false,
	Version:       version.Version,
	Use:           "serenity <command> [flags]",
	Short:         "Serenity is an aggressive, no-noise and ultra fast Go linter",
}

func Exec() {
	err := rootCmd.Execute()

	os.Exit(exception.ExitCode(err))
}

func init() {
	rootCmd.PersistentFlags().Bool("no-color", false, "Remove color from the output")
	rootCmd.PersistentFlags().Bool("verbose", false, "Print additional diagnostics and processed files")
	rootCmd.PersistentFlags().String("config", "", "Path to configuration file (Auto-discovered if omitted)")
}
