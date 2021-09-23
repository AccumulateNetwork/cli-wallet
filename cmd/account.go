package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Create and get token accounts",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			switch arg := args[0]; arg {
			case "get":
				if len(args) > 1 {
					GetAccount(args[1])
				} else {
					fmt.Println("Usage:")
					PrintAccountGet()
				}
			case "create":
				if len(args) > 3 {
					CreateAccount(args[1], args[2], args[3])
				} else {
					fmt.Println("Usage:")
					PrintAccountCreate()
				}
			default:
				fmt.Println("Usage:")
				PrintAccount()
			}
		} else {
			fmt.Println("Usage:")
			PrintAccount()
		}

	},
}

func init() {
	rootCmd.AddCommand(accountCmd)
}

func PrintAccountGet() {
	fmt.Println("  accumulate account get [URL]					Get token account by URL")
}

func PrintAccountCreate() {
	fmt.Println("  accumulate account create [URL] [TOKEN URL] [SIGNER ADI]	Create new token account")
}

func PrintAccount() {
	PrintAccountGet()
	PrintAccountCreate()
}

func GetAccount(url string) {
	fmt.Println("Getting token account " + url)
}

func CreateAccount(url string, token string, signer string) {
	fmt.Println("Creating new token account " + url)
}
