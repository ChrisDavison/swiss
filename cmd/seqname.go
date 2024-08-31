/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// seqnameCmd represents the seqname command
var seqnameCmd = &cobra.Command{
	Use:   "seqname",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("seqname called")
	},
}

func init() {
	rootCmd.AddCommand(seqnameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// seqnameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// seqnameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
