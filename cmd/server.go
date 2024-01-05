/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ahelmy/xdev/internal"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "Start a web server, Alias: s",
	Long: `Start a web server. For example:
	server
	server -p 8000`,
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetInt32("port")
		if err != nil || port != 0 {
			port = 8000
		}
		internal.StartServer(port)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().Int32P("port", "p", 8000, "Server port")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
