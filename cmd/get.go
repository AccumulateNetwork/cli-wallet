package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get data by URL",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			Get(args[0])
		} else {
			fmt.Println("Usage:")
			PrintGet()
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func PrintGet() {
	fmt.Println("  accumulate get [URL] 		Get data by Accumulate URL")
}

func Get(url string) {
	fmt.Println("Getting data by Accumulate URL: " + url)
}
