package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Issue and get tokens",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			switch arg := args[0]; arg {
			case "get":
				if len(args) > 1 {
					GetToken(args[1])
				} else {
					fmt.Println("Usage:")
					PrintTokenGet()
				}
			case "create":
				if len(args) > 4 {
					CreateToken(args[1], args[2], args[3], args[4])
				} else {
					fmt.Println("Usage:")
					PrintTokenCreate()
				}
			default:
				fmt.Println("Usage:")
				PrintToken()
			}
		} else {
			fmt.Println("Usage:")
			PrintToken()
		}

	},
}

func init() {
	rootCmd.AddCommand(tokenCmd)
}

func PrintTokenGet() {
	fmt.Println("  accumulate token get [URL]						Get token by URL")
}

func PrintTokenCreate() {
	fmt.Println("  accumulate token create [URL] [SYMBOL] [PRECISION] [SIGNER ADI]	Create new token")
}

func PrintToken() {
	PrintTokenGet()
	PrintTokenCreate()
}

func GetToken(url string) {
	fmt.Println("Getting token " + url)
}

func CreateToken(url string, symbol string, precision string, signer string) {
	fmt.Println("Creating new token " + symbol)
}
