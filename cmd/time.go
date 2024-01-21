/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/ahelmy/xdev/internal"
	"github.com/spf13/cobra"
)

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:     "time",
	Aliases: []string{"t"},
	Short:   "Time command to covert and manupulate time",
	Long: `Covert time and dates to different formats and timezones. 
	Formats:
		d 	Day of the month as digits; no leading zero for single-digit days.
		M 	Month as digits; no leading zero for single-digit months.
		yy 	Year as last two digits; leading zero for years less than 10.
		H 	Hour; no leading zero for single-digit hours (24-hour clock).
		m 	Minute; no leading zero for single-digit minutes.
		s 	Second; no leading zero for single-digit seconds.
		Z 	ISO 8601 offset from UTC in timezone (+hhmm), or Z if UTC.
	For example:
	time
	UTC: 20-01-2024 13:30:19 UTC
	Your Timezone: 20-01-2024 15:30:19 EET
	Epoch: 1705757419

	time -f "d M yy"
	UTC: 20 1 24
	Your Timezone: 20 1 24
	Epoch: 1705757438

	time --from epoch 1705751771
	UTC: 20-01-2024 11:56:11 UTC
	Your Timezone: 20-01-2024 13:56:11 EET
	Epoch: 1705751771

	time --from dd-MM-yyyy -f dd-yy 24-01-2024
	UTC: 24-24
	Your Timezone: 24-24
	Epoch: 1706054400`,
	Run: func(cmd *cobra.Command, args []string) {
		format := internal.ParseFormat(cmd.Flag("format").Value.String())
		from := cmd.Flag("from").Value.String()
		if len(args) == 0 {
			fmt.Println(internal.Now(format))
			return
		}
		if from == "epoch" {
			epoch, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(internal.ConvertTimeFromEpoch(epoch, internal.ToDateFormat))
		} else {
			if from == "" {
				fmt.Println("Please specify the from format")
				return
			}
			time, err := internal.ConvertTimeFromFormat(args[0], from, internal.ToDateFormat)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(time)
		}
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)
	timeCmd.Flags().StringP("from", "r", "", "From type. epoch or format")
	timeCmd.Flags().StringP("format", "f", internal.DateFormat, `Format to convert to.
		d 	Day of the month as digits; no leading zero for single-digit days.
		M 	Month as digits; no leading zero for single-digit months.
		yy 	Year as last two digits; leading zero for years less than 10.
		H 	Hour; no leading zero for single-digit hours (24-hour clock).
		m 	Minute; no leading zero for single-digit minutes.
		s 	Second; no leading zero for single-digit seconds.
		Z 	ISO 8601 offset from UTC in timezone (+hhmm), or Z if UTC.`)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
