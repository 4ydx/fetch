package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/net/html"
)

type Meta struct {
	Site       string
	LinkCount  int
	ImageCount int
	LastFetch  time.Time
}

func getMeta(url string) (Meta, error) {
	meta := Meta{Site: url}

	dir := directory(url)
	file, err := os.Open(fmt.Sprintf("%s/index.html", dir))
	if err != nil {
		return meta, err
	}

	info, err := file.Stat()
	if err != nil {
		return meta, err
	}
	meta.LastFetch = info.ModTime()

	tokens := html.NewTokenizer(file)
	for {
		switch tokens.Next() {
		case html.ErrorToken:
			return meta, nil
		case html.StartTagToken:
			token := tokens.Token()

			if token.Data == "a" {
				meta.LinkCount++
			}
			if token.Data == "img" {
				meta.ImageCount++
			}
		}
	}
}
