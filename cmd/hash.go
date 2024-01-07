/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/ahelmy/xdev/internal"
	"github.com/spf13/cobra"
)

var validAlgorithms = []string{"md5", "sha256", "sha512", "salt"}

func validateAlgorithm(algorithm string) bool {
	for _, a := range validAlgorithms {
		if a == algorithm {
			return true
		}
	}
	return false
}

// hashCmd represents the hash command
var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Hash a string",
	Long: `Hash a string. For example:
	hash -a md5 "abc"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide an algorithm and a string to hash")
			return
		}
		algorithm, err := cmd.Flags().GetString("algorithm")
		if err != nil {
			algorithm = "md5"
		}
		if !validateAlgorithm(algorithm) {
			fmt.Println("Please provide a valid algorithm, valid algorithms are: " + strings.Join(validAlgorithms, ", "))
			return
		}
		result, err := internal.HashString(args[0], algorithm)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(hashCmd)
	hashCmd.Flags().StringP("algorithm", "a", "md5", "Hash algorithm")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hashCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hashCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
