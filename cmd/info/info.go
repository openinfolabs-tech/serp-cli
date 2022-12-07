/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"fmt"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("info called")
	},
}
