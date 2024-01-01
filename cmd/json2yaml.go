/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ahelmy/xdev/internal"
)

// json2yamlCmd represents the json2yaml command
var json2yamlCmd = &cobra.Command{
	Use:     "json2yaml",
	Aliases: []string{"j2y"},
	Short:   "Convert JSON to YAML. Alias: j2y",
	Long: `Convert JSON to YAML. For example:
	
	j2y '{"name":"ahmed","age":30}'
	Output:
	name: ahmed
	age: 30`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a JSON string.")
			return
		}
		yaml, err := internal.Json2Yaml(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(yaml)
	},
}

func init() {
	rootCmd.AddCommand(json2yamlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// json2yamlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// json2yamlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
