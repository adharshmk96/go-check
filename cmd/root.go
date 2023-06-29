package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "go-check",
	Short: "A CLI app to find all the unique import statements within a folder for a Golang application",
	Long: `go-check is a CLI application that allows to find all the unique import statements 
within a folder for a Golang application by providing the path to the directory as an argument.`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
