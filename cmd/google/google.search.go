/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package google

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var (
	query    string
	pages    string
	output   string
	savePath string
)

func trimLeftChars(s string, n int) string {
	m := 0
	for i := range s {
		if m >= n {
			return s[i:]
		}
		m++
	}
	return s[:0]
}

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

	var jsonResults []interface{}

	var paginationIndex = 0
	totalPages, err := strconv.Atoi(pages)
	if err != nil {
		panic(err)
	}

	var initialUrl string = fmt.Sprintf("https://www.google.com/search?q=%s&client=firefox-b-e", url.QueryEscape(searchQuery))
	var nextPage string = ""

	// Create a new collector
	c := colly.NewCollector()

	if err != nil {
		log.Fatal(err)
	}
	c.SetRequestTimeout(60 * time.Second)
	q, _ := queue.New(
		2, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)

	c.OnRequest(func(r *colly.Request) {
		setHeaders(r)
	})

	c.OnResponse(func(r *colly.Response) {
		// r.Save(fmt.Sprintf("%d.html", paginationIndex))
		paginationIndex += 1
	})

	// Set HTML callback for pagination
	c.OnHTML("#pnnext", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if paginationIndex < totalPages {
			fmt.Println("Loading next page: ", fmt.Sprintf("https://google.com%s&client=firefox-b-e", link))
			q.AddURL(fmt.Sprintf("https://google.com/%sclient=firefox-b-e", link))
		}
	})

	c.OnHTML("#cnt", func(e *colly.HTMLElement) {
		// for each search engine result
		e.ForEach(".MjjYud", func(_ int, el *colly.HTMLElement) {
			var breadcrumb string = el.ChildText("div.TbwUpd.NJjxre cite")
			var heading string = el.ChildText("a h3.LC20lb.MBeuO.DKV0Md")
			var urlString string = el.ChildAttr("div.yuRUbf a", "href")
			var description string = el.ChildText("div.VwiC3b.yXK7lf.MUxGbd.yDYNvb.lyLwlc.lEBKkf")
			var timeAgo = el.ChildText("span.MUxGbd.wuQ4Ob.WZ8Tjf")
			var jsonObj map[string]interface{}

			err := json.Unmarshal([]byte("{}"), &jsonObj)
			if err != nil {
				fmt.Println(err)
				return
			}

			if len(heading) > 0 && len(urlString) > 0 && len(description) > 0 {
				if output == "tui" {
					fmt.Println("")
					fmt.Printf("%s\n", aurora.Magenta(heading))
					fmt.Println(aurora.White(description))
					// fmt.Printf("%s", aurora.Gray(20-1, breadcrumb))
					fmt.Printf("%s\n", aurora.Cyan(urlString))
				}
				if output == "json" {
					jsonObj["heading"] = heading
					jsonObj["description"] = description
					jsonObj["url"] = urlString
					jsonObj["breadcrumb"] = breadcrumb
					jsonObj["timeAgo"] = timeAgo
					jsonResults = append(jsonResults, jsonObj)

				}
			}
		})

		jsonArrVal, _ := json.Marshal(jsonResults)

		if savePath != "" {
			if paginationIndex >= totalPages {
				f, err := os.Create(savePath)
				if err != nil {
					fmt.Println(err)
					return
				}
				w := bufio.NewWriter(f)
				totalBytes, err := w.WriteString(string(jsonArrVal))
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Printf("wrote %d bytes to %s\n", totalBytes, savePath)
				w.Flush()
			}

		} else {
			// else no save path set, log results to console
			fmt.Println(string(jsonArrVal))
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
	googleSearchCmd.Flags().StringVarP(&savePath, "file", "f", "", "specify the path where results will be saved")
	googleSearchCmd.Flags().StringVarP(&output, "output", "o", "", "specify the output format")
	googleSearchCmd.Flags().StringVarP(&query, "query", "q", "", "The google search query")
	googleSearchCmd.Flags().StringVarP(&pages, "pages", "p", "1", "Total number of pages to scrape, default is 1 page")

	if err := googleSearchCmd.MarkFlagRequired("query"); err != nil {
		fmt.Println(err)
	}
	if err := googleSearchCmd.MarkFlagRequired("output"); err != nil {
		fmt.Println(err)
	}

	GoogleCmd.AddCommand(googleSearchCmd)
}
