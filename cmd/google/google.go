/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package google

import (
	"github.com/spf13/cobra"
)

// googleCmd represents the google command
var GoogleCmd = &cobra.Command{
	Use:   "google",
	Short: "crawl google.com (--query, -q), (--pages,-p)",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
