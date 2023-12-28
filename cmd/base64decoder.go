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
var base64decoderCmd = &cobra.Command{
	Use:     "base64decoder",
	Aliases: []string{"b64d"},
	Short:   "Decode base64 string. Alias: b64d",
	Long: `Decode a base64 string. For example:

base64decoder "YWJjMTIzIT8kKiYoKSctPUB+"
Output: abc123!?$*&()'-=@~`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a base64 string to decode")
			return
		}
		fmt.Println(internal.DecodeFromBase64(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(base64decoderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// base64encoderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// base64encoderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
