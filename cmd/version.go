/*
Copyright © 2023 Adharsh M adharshmk96@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "display the version of " + APP_NAME,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(APP_NAME+" version: %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
