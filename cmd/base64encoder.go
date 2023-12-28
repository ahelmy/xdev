/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	internal "github.com/ahelmy/xdev/internal"
	"github.com/spf13/cobra"
)

// base64encoderCmd represents the base64encoder command
var base64encoderCmd = &cobra.Command{
	Use:     "base64encoder",
	Aliases: []string{"b64e"},
	Short:   "Encode string to base64. Alias: b64e",
	Long: `Encode a string to a base 64 string. For example:

base64encoder "abc123!?$*&()'-=@~"
Output: YWJjMTIzIT8kKiYoKSctPUB+`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a string to encode")
			return
		}
		fmt.Println(internal.EncodeToBase64(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(base64encoderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// base64encoderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// base64encoderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
