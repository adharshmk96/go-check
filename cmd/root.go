package cmd

import "github.com/spf13/cobra"

var APP_NAME = "go-check"
var version = "v0.0.0"

var rootCmd = &cobra.Command{
	Use:   "go-check",
	Short: "A CLI app to verify rules and inspect the codebase of a golang project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
