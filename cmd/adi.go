package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var adiCmd = &cobra.Command{
	Use:   "adi",
	Short: "Create and manage ADI",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			switch arg := args[0]; arg {
			case "get":
				if len(args) > 1 {
					GetADI(args[1])
				} else {
					fmt.Println("Usage:")
					PrintADIGet()
				}
			case "create":
				if len(args) > 1 {
					NewADI(args[1])
				} else {
					fmt.Println("Usage:")
					PrintADICreate()
				}
			case "import":
				if len(args) > 2 {
					ImportADI(args[1], args[2])
				} else {
					fmt.Println("Usage:")
					PrintADIImport()
				}
			default:
				fmt.Println("Usage:")
				PrintADI()
			}
		} else {
			fmt.Println("Usage:")
			PrintADI()
		}

	},
}

func init() {
	rootCmd.AddCommand(adiCmd)
}

func PrintADIGet() {
	fmt.Println("  accumulate adi get [URL]			Get existing ADI from database")
}

func PrintADICreate() {
	fmt.Println("  accumulate adi create [URL] [SIGNER ADI]	Create new ADI")
}

func PrintADIImport() {
	fmt.Println("  accumulate adi import [URL] [PRIVATE KEY]	Import Existing ADI")
}

func PrintADI() {
	PrintADIGet()
	PrintADICreate()
	PrintADIImport()
}

func GetADI(url string) {
	fmt.Println("Getting ADI " + url + " from local DB")
}

func NewADI(url string) {
	fmt.Println("Generating new ADI " + url)
}

func ImportADI(url string, pk string) {
	fmt.Println("Importing ADI " + url + " with private key " + pk)
}
