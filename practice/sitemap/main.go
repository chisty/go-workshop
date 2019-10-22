package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/chisty/customhtmlutil"
)

type empty struct{}

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "url for sitemap")
	maxDepth := flag.Int("depth", 3, "maximum depth of website")
	flag.Parse()

	// links := getPage(*urlFlag)
	// for _, l := range links {
	// 	fmt.Println(l)
	// }

	pages := bfs(*urlFlag, *maxDepth)
	for _, page := range pages {
		fmt.Println(page)
	}
}

func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]empty)
	queue := make(map[string]empty)
	nextQueue := map[string]empty{urlStr: empty{}}

	for i := 0; i <= maxDepth; i++ {
		queue, nextQueue = nextQueue, make(map[string]empty)
		for url, _ := range queue {
			if _, exists := seen[url]; exists == false {
				seen[url] = struct{}{}
				for _, link := range getPage(url) {
					nextQueue[link] = empty{}
				}
			}
		}
	}

	result := make([]string, 0, len(seen))
	for url, _ := range seen {
		result = append(result, url)
	}
	return result
}

func getPage(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()

	return filter(getLinks(resp.Body, base), withPrefix(base))
}

func getLinks(reader io.Reader, baseURL string) []string {
	var linkResult []string
	links, _ := customhtmlutil.Parse(reader)

	for _, link := range links {
		switch {
		case strings.HasPrefix(link.Href, "/"):
			linkResult = append(linkResult, baseURL+link.Href)
		case strings.HasPrefix(link.Href, "http"):
			linkResult = append(linkResult, link.Href)
		}
	}

	return linkResult
}

func filter(links []string, subFilterFunc func(string) bool) []string {
	var result []string
	for _, link := range links {
		if subFilterFunc(link) {
			result = append(result, link)
		}
	}
	return result
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
