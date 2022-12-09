/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package google

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/spf13/cobra"
)

var (
	query string
	pages string
)

func setHeaders(r *colly.Request) {
	r.Headers.Set("Host", "www.google.com")
	r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:102.0) Gecko/20100101 Firefox/102.0")
	r.Headers.Set("Accept", "text/html")
	r.Headers.Set("Accept-Language", "en-US,en;q=0.5")
	r.Headers.Set("Accept-Encoding", "text/html")
	r.Headers.Set("Connection", "keep-alive")
	r.Headers.Set("Upgrade-Insecure-Requests", "1")
	r.Headers.Set("Sec-Fetch-Dest", "document")
	r.Headers.Set("Sec-Fetch-Mode", "navigate")
	r.Headers.Set("Sec-Fetch-Site", "same-site")
	r.Headers.Set("TE", "trailers")
}

func crawlGoogle(searchQuery string) {
	var paginationIndex = 0
	totalPages, err := strconv.Atoi(pages)
	if err != nil {
		panic(err)
	}

	var initialUrl string = fmt.Sprintf("https://www.google.com/search?q=%s&client=firefox-b-e", searchQuery)
	var nextPage string = ""

	// Create a new collector
	c := colly.NewCollector()

	if err != nil {
		log.Fatal(err)
	}
	c.SetRequestTimeout(60 * time.Second)
	q, _ := queue.New(
		1, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)

	c.OnRequest(func(r *colly.Request) {
		setHeaders(r)
	})

	c.OnResponse(func(r *colly.Response) {
		r.Save(fmt.Sprintf("%d.html", paginationIndex))
		paginationIndex += 1
	})

	// Set HTML callback
	c.OnHTML("#pnnext", func(e *colly.HTMLElement) {
		// "a.nBDE1b:nth-child(3)"
		link := e.Attr("href")
		if paginationIndex < totalPages {
			fmt.Println("Loading next page: ", fmt.Sprintf("https://google.com%s&client=firefox-b-e", link))
			q.AddURL(fmt.Sprintf("https://google.com%sclient=firefox-b-e", link))
		}
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	if nextPage == "" {
		fmt.Println("Loading first page: ", initialUrl)
		q.AddURL(initialUrl)
	}
	q.Run(c)
}

// infoCmd represents the info command
var googleSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		msg := fmt.Sprintf("Starting crawl job for google.com, query: %s", query)
		fmt.Println(msg)
		crawlGoogle(query)
	},
}

func init() {
	googleSearchCmd.Flags().StringVarP(&query, "query", "q", "", "The google search query")
	googleSearchCmd.Flags().StringVarP(&pages, "pages", "p", "", "Total number of pages to scrape, default is 1 page")

	if err := googleSearchCmd.MarkFlagRequired("query"); err != nil {
		fmt.Println(err)
	}

	GoogleCmd.AddCommand(googleSearchCmd)
}
