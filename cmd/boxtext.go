/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// boxtextCmd represents the boxtext command
var boxtextCmd = &cobra.Command{
	Use:   "boxtext [flags] text...",
	Short: "Wrap text in ascii box drawing",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		text := strings.Join(args, " ")
		dashes := strings.Repeat("─", len(text)+2)

		fmt.Printf("┌%v┐\n", dashes)
		fmt.Printf("│ %v │\n", text)
		fmt.Printf("└%v┘", dashes)
	},
}

func init() {
	rootCmd.AddCommand(boxtextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// boxtextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// boxtextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
