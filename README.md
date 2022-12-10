# serp-search

A work-in-progress Google and Google CSE CLI crawler...

`./toolbox google search -p 2 -o json -f google-results.json --query "how to cast to int js" `

`./toolbox google search -p 2 --output tui -q "how to cast to int js"`


```bash
./toolbox google search -h
# Usage:
#   toolbox google search [flags]

# Flags:
#   -h, --help           help for search
#   -p, --pages string   Total number of pages to scrape, default is 1 page
#   -q, --query string   The google search query
```

![demo](./docs/demo.gif)