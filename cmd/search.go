/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search note by string or date created",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
