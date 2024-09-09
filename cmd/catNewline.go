/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// catNewlineCmd represents the catNewline command
var catNewlineCmd = &cobra.Command{
	Use:   "catNewline",
	Short: "Join files together with newline in between each",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("catNewline called")
	},
}

func init() {
	rootCmd.AddCommand(catNewlineCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// catNewlineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// catNewlineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
