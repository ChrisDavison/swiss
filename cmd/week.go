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

// weekCmd represents the week command
var weekCmd = &cobra.Command{
	Use:   "week",
	Short: "This weeks' previous Monday and future Sunday date",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		today := time.Now()
		shift := int(today.Weekday()-1) * 24
		durn, err := time.ParseDuration(fmt.Sprintf("%dh", shift))
		if err != nil {
			log.Fatal(err)
		}
		startOfWeek := today.Add(-durn)
		endOfWeek := startOfWeek.Add(6 * 24 * time.Hour)
		fmt.Printf("%s to %s\n", startOfWeek.Format("2006-01-02"), endOfWeek.Format("2006-01-02"))

	},
}

func init() {
	rootCmd.AddCommand(weekCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// weekCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// weekCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
