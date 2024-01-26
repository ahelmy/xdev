package cmd

import (
	"fmt"
	"github.com/ahelmy/xdev/internal"
	"github.com/spf13/cobra"
)

var properties2yamlCmd = &cobra.Command{
	Use:   "p2y",
	Short: "Convert Java Properties to YAML",
	Long:  `Convert Java Properties to YAML`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a Java Properties string.")
			return
		}
		res, err := internal.Properties2Yaml(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(properties2yamlCmd)
}
