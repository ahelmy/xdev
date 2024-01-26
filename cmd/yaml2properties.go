package cmd

import (
	"fmt"
	"github.com/ahelmy/xdev/internal"
	"github.com/spf13/cobra"
)

var yaml2propertiesCmd = &cobra.Command{
	Use:   "yaml2properties",
	Aliases: []string{"y2p"},
	Short: "Convert YAML to Java Properties, Alias: y2p",
	Long:  `Convert YAML to Java Properties`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a YAML string.")
			return
		}
		res, err := internal.Yaml2Properties(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(yaml2propertiesCmd)
}
