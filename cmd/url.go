/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ahelmy/xdev/internal"
	"github.com/spf13/cobra"
)

// urlCmd represents the url command
var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "URL encode/decode.",
	Long: `URL encode/decode. For example:
	encode a URL string (default).
	decode a URL string.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}
	},
}

var urlEncodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "URL encode a string.",
	Long: `URL encode a string. For example:
	encode "https://www.google.com/search?q=hello+world"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a URL string to encode")
			return
		}
		fmt.Println(internal.EncodeURL(args[0]))
	},
}

var urlDecodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "URL decode a string.",
	Long: `URL decode a string. For example:
	decode "https%3A%2F%2Fwww.google.com%2Fsearch%3Fq%3Dhello%2Bworld"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a URL string to decode")
			return
		}
		decoded, err := internal.DecodeURL(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(decoded)
	},
}

func init() {
	rootCmd.AddCommand(urlCmd)
	urlCmd.AddCommand(urlEncodeCmd, urlDecodeCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// urlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// urlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
