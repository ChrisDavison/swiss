/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/spf13/cobra"
)

// kgForBmiCmd represents the kgForBmi command
var kgForBmiCmd = &cobra.Command{
	Use:   "kgForBmi KG",
	Args:  cobra.ExactArgs(1),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		kg, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("BMI %.0f is %.2f Kg\n", kg, kg*math.Pow(1.78, 2))
	},
}

func init() {
	rootCmd.AddCommand(kgForBmiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kgForBmiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kgForBmiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
