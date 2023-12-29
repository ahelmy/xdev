/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	internal "github.com/ahelmy/xdev/internal"
	"github.com/spf13/cobra"
)

// paswordCmd represents the pasword command
var passwordCmd = &cobra.Command{
	Use:     "password",
	Aliases: []string{"pwd"},
	Short:   "Password generator. Alias: pwd",
	Long: `Password generator. For example:
	
	pwd
	Output: 5b3e4f2a1#
	
	pwd -l 20
	Output: 5b@e4F2a1#qqwweerrt
	
	pwd -l 20 -e=false
	Output: 5b3e4F2a12qqwweerrt`,
	Run: func(cmd *cobra.Command, args []string) {
		lengthStr := cmd.Flag("length").Value.String()
		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			length = 10
		}
		espcial, err := cmd.Flags().GetBool("espcial")
		if err != nil {
			espcial = true
		}
		numeric, err := cmd.Flags().GetBool("numeric")
		if err != nil {
			numeric = true
		}
		capital, err := cmd.Flags().GetBool("capital")
		if err != nil {
			capital = true
		}
		fmt.Println(internal.GeneratePassword(length, espcial, numeric, capital))
	},
}

func init() {
	rootCmd.AddCommand(passwordCmd)
	passwordCmd.Flags().Int32P("length", "l", 10, "Password length")
	passwordCmd.Flags().BoolP("espcial", "e", true, "Especial characters")
	passwordCmd.Flags().BoolP("numeric", "n", true, "Numeric characters")
	passwordCmd.Flags().BoolP("capital", "c", true, "Capital characters")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// paswordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// paswordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
