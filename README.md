# serp-search

A work-in-progress Google and Google CSE CLI crawler...

`./serp google search -p 2 -o json -f google-results.json --query "how to cast to int js" `

`./serp google search -p 2 --output tui -q "how to cast to int js"`

Or pipe into something like jq :
`./serp google search -p 10 -q 'intext:"index of" ".sql"' | jq`

```bash
./serp google search -h
# Usage:
#   serp google search [flags]

# Flags:
#   -f, --file string     specify the path where results will be saved
#   -h, --help            help for search
#   -o, --output string   specify the output format (json,tui) (default "json")
#   -p, --pages string    Total number of pages to scrape (default "1")
#   -q, --query string    The google search query
```

[![asciicast](https://user-images.githubusercontent.com/29207058/208087767-feefe329-30ab-45b1-9526-004aa79f63a2.gif)](https://user-images.githubusercontent.com/29207058/208087767-feefe329-30ab-45b1-9526-004aa79f63a2.gif)

# License
MIT
