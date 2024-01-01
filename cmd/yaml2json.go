package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ahelmy/xdev/internal"
)

// yaml2jsonCmd represents the yaml2json command
var yaml2jsonCmd = &cobra.Command{
	Use:     "yaml2json",
	Aliases: []string{"y2j"},
	Short:   "Convert YAML to JSON. Alias: y2j",
	Long: `Convert YAML to JSON. For example:

	y2j 'name: ahmed
	age: 30'
	Output:
	{"name":"ahmed","age":30}`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a YAML string.")
			return
		}
		jsonString, err := internal.Yaml2Json(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(jsonString)
	},
}

func init() {
	rootCmd.AddCommand(yaml2jsonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// json2yamlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// json2yamlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
