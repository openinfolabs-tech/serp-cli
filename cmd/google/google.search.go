/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package google

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
)

var (
	query string
)

func crawlGoogle() {
	// Create a new collector
	c := colly.NewCollector()

	// Set HTML callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println(link)
	})

	c.OnResponse()

	// Start scraping on https://en.wikipedia.org
	var initialUrl = fmt.Sprintf("https://www.google.com/search?q=%s", query)  
	c.Visit(initialUrl)
}

// infoCmd represents the info command
var googleSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		msg := fmt.Sprintf("Starting crawl job for google.com, query: %s", query)
		fmt.Println(msg)
		crawlGoogle()
	},
}

func init() {

	googleSearchCmd.Flags().StringVarP(&query, "query", "q", "", "The google search query")

	if err := googleSearchCmd.MarkFlagRequired("query"); err != nil {
		fmt.Println(err)
	}

	GoogleCmd.AddCommand(googleSearchCmd)
}
