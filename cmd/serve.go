/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var port int
var directory string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Basic static file server",
	Run: func(cmd *cobra.Command, args []string) {
		http.Handle("/", http.FileServer(http.Dir(directory)))

		log.Printf("serving `%s` on :%d\n", directory, port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&port, "port", "p", 8000, "Port")
	serveCmd.Flags().StringVarP(&directory, "directory", "d", ".", "Directory")
}
