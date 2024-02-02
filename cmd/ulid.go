/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ahelmy/xdev/internal"
	"github.com/spf13/cobra"
)

// ulidCmd represents the ulid command
var ulidCmd = &cobra.Command{
	Use:   "ulid",
	Short: "Generate a ulid string",
	Long:  `Generate a ulid string. For example:`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(internal.GenerateULID())
		} else {
			to := cmd.Flag("to").Value.String()
			switch to {
			case "uuid":
				ulid, err := internal.ULIDtoUUID(args[0])
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
	rootCmd.AddCommand(ulidCmd)
	ulidCmd.Flags().StringP("to", "t", "ulid", "Convert from ulid to (uuid)")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ulidCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ulidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
