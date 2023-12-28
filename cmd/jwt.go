/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	internal "github.com/ahelmy/xdev/internal"
	"github.com/spf13/cobra"
)

// jwtdecodeCmd represents the jwtdecode command
var jwtCmd = &cobra.Command{
	Use:     "jwt",
	Aliases: []string{"jwt"},
	Short:   "Decode or encode a JWT string.",
	Long: `Decode or encode a JWT string. For example:
	
	jwt "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJBaGVsbXkiLCJleHAiOjE1NjY5NjQwMzB9.2ZQ5.
	jwt {header} {payload} {secret}`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a JWT string to decode")
			return
		}
		if args[0][0] == 'e' {
			jwt, err := internal.DecodeJWT(args[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Header: \n", jwt.Header)
			fmt.Println("Payload: \n", jwt.Claims)
			fmt.Println("Signature: \n", jwt.Signature)
		} else {
			fmt.Println("Not implemented yet")
		}
	},
}

func init() {
	rootCmd.AddCommand(jwtCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jwtdecodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jwtdecodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
