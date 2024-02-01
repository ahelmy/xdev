/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ahelmy/xdev/internal"
)

// guidCmd represents the guid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate a uuid string",
	Long:  `Generate a uuid string.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(internal.GenerateUUID())
		} else {
			to := cmd.Flag("to").Value.String()
			switch to {
			case "ulid":
				ulid, err := internal.UUIDtoULID(args[0])
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(ulid)
				}
			default:
				fmt.Println("Unknown conversion")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)
	uuidCmd.Flags().StringP("to", "t", "ulid", "Convert from uuid to (ulid)")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// guidCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// guidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
