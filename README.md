# serp-search

A work-in-progress Google and Google CSE CLI crawler...

`./toolbox google search -p 2 -o json -f google-results.json --query "how to cast to int js" `

`./toolbox google search -p 2 --output tui -q "how to cast to int js"`


```bash
./toolbox google search -h
# Usage:
#   toolbox google search [flags]

# Flags:
#   -f, --file string     specify the path where results will be saved
#   -h, --help            help for search
#   -o, --output string   specify the output format (json,tui) (default "json")
#   -p, --pages string    Total number of pages to scrape (default "1")
#   -q, --query string    The google search query
```

![demo](https://user-images.githubusercontent.com/29207058/206865617-22da83f2-e74b-4b75-9116-a8874811dde3.gif)


# License
MIT