/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	internal "github.com/ahelmy/xdev/internal"
	"github.com/spf13/cobra"
)

// guidCmd represents the guid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate a uuid string",
	Long:  `Generate a uuid string.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.GenerateGUID())
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// guidCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// guidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
