package cmd

import (
	"fmt"

	"github.com/adharshmk96/go-check/pkg"
	"github.com/spf13/cobra"
)

var path string
var pattern string

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View all unique import statements",
	Long:  `View all unique import statements within a folder for a Golang application by providing the path to the directory as an argument.`,
	Run: func(cmd *cobra.Command, args []string) {
		if path == "" {
			fmt.Println("Please provide the path to the directory as -p flag")
			return
		}
		files, err := pkg.GetGoFiles(path)
		if err != nil {
			fmt.Println(err)
			return
		}
		imports, err := pkg.GetImports(files)
		if err != nil {
			fmt.Println(err)
			return
		}
		pkg.DisplayImports(imports, pattern)
	},
}

func init() {
	viewCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "Path to the directory")
	viewCmd.PersistentFlags().StringVarP(&pattern, "pattern", "i", "", "Import starts with")
	rootCmd.AddCommand(viewCmd)
}
