/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	internal "github.com/ahelmy/xdev/internal"
	"github.com/spf13/cobra"
)

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "JSON indentation and minification",
	Long: `JSON indentation and minification. For example:
	json '{"name":"ahmed","age":30}'
	Output: {
		"name": "ahmed",
		"age": 30
		}
	json -m {
		"name": "ahmed",
		"age": 30
		}
	Output: {"name":"ahmed","age":30}`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a JSON string.")
			return
		}
		if cmd.Flag("minify").Value.String() == "true" {
			minify, err := internal.MinifyJSON(args[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(minify)
		} else {
			indent, err := internal.IndentJSON(args[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(indent)
		}
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)
	jsonCmd.Flags().BoolP("minify", "m", false, "Minify the JSON string")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jsonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
