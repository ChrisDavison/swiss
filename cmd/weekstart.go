/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

// weekstartCmd represents the weekstart command
var weekstartCmd = &cobra.Command{
	Use:   "weekstart",
	Short: "Print date of this week's Monday",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		today := time.Now()
		shift := int(today.Weekday()-1) * 24
		durn, err := time.ParseDuration(fmt.Sprintf("%dh", shift))
		if err != nil {
			log.Fatal(err)
		}
		startOfWeek := today.Add(-durn)
		fmt.Println(startOfWeek.Format("2006-01-02"))
	},
}

func init() {
	rootCmd.AddCommand(weekstartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// weekstartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// weekstartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
