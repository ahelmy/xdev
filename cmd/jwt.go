/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
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
	jwt -h {"alg": "HS256"} -c '{"name":"ali"}' -s 123`,
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
			headersStr := cmd.Flag("headers").Value.String()
			headers := make(map[string]interface{})
			err := json.Unmarshal([]byte(headersStr), &headers)
			if err != nil {
				fmt.Println("Invalid headers format:", err)
				return
			}
			claimsStr := cmd.Flag("claims").Value.String()
			claims := make(map[string]interface{})
			err = json.Unmarshal([]byte(claimsStr), &claims)
			if err != nil {
				fmt.Println("Invalid claims format:", err)
				return
			}
			secret := cmd.Flag("secret").Value.String()
			jwt, err := internal.EncodeJWT(headers, claims, secret)
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
	jwtCmd.Flags().StringP("headers", "d", "{}", "The header to encode in the JWT")
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
