/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var headingLevel int

// setexCmd represents the setex command
var setexCmd = &cobra.Command{
	Use:   "setex [flags] text...",
	Short: "A brief description of your command",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		text := strings.Join(args, " ")
		switch headingLevel {
		case 1:
			underline := strings.Repeat("=", len(text))
			fmt.Printf("%s\n%s\n", text, underline)
		case 2:
			underline := strings.Repeat("-", len(text))
			fmt.Printf("%s\n%s\n", text, underline)
		default:
			markers := strings.Repeat("#", headingLevel)
			fmt.Printf("%s %s", markers, text)
		}
	},
}

func init() {
	rootCmd.AddCommand(setexCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setexCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	setexCmd.Flags().IntVarP(&headingLevel, "level", "l", 1, "Heading level")
}
