package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	meta := flag.String("metadata", "", "print metadata for downloaded url")
	flag.Parse()

	if meta != nil && *meta != "" {
		m, err := getMeta(*meta)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			return
		}
		fmt.Printf(`site: %s
num_links: %d
images: %d
last_fetch: %s
`, m.Site, m.LinkCount, m.ImageCount, m.LastFetch.Format(time.UnixDate))
		return
	}

	urls := flag.Args()
	for _, url := range urls {
		err := download(url)
		if err != nil {
			fmt.Printf("%s %s\n", url, err)
		}
	}
}
