/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ahelmy/xdev/internal"
)

// jwtdecodeCmd represents the jwtdecode command
var jwtCmd = &cobra.Command{
	Use:     "jwt",
	Aliases: []string{"jwt"},
	Short:   "Decode or encode a JWT string.",
	Long: `Decode or encode a JWT string. For example:
	
	jwt "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJBaGVsbXkiLCJleHAiOjE1NjY5NjQwMzB9.2ZQ5"
	jwt -a HS256 -c '{"name":"ali"}' -s 123`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 && args[0][0] == 'e' {
			jwt, err := internal.DecodeJWT(args[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Header:\n", jwt.Header)
			fmt.Println("Payload:\n", jwt.Claims)
			expiry := "Never"
			if jwt.Expires != nil {
				expiry = jwt.Expires.String()
			}
			fmt.Println("Expires at: ", expiry)
		} else {
			algorithm := cmd.Flag("algorithm").Value.String()
			claims := cmd.Flag("claims").Value.String()
			secret := cmd.Flag("secret").Value.String()
			jwt, err := internal.EncodeJWT(algorithm, claims, secret)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(jwt)
		}
	},
}

func init() {
	rootCmd.AddCommand(jwtCmd)
	jwtCmd.Flags().StringP("algorithm", "a", "HS256", "The algorithm to use for signing the JWT")
	jwtCmd.Flags().StringP("claims", "c", "{}", "The claims to encode in the JWT")
	jwtCmd.Flags().StringP("secret", "s", "", "The secret to use for signing the JWT")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jwtdecodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jwtdecodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
