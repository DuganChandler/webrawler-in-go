package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	htmlReader := strings.NewReader(htmlBody)
	htmlParse, err := html.Parse(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("unable to parse html: %v", err)
	}

	var urls []string
	htmlParseHelper(htmlParse, baseURL, &urls)

	return urls, nil
}

func htmlParseHelper(n *html.Node, baseUrl *url.URL, urls *[]string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				href, err := url.Parse(a.Val)
				if err != nil {
					fmt.Printf("could not parse href '%v': %v\n", a.Val, err)
					continue
				}

				resolvedURL := baseUrl.ResolveReference(href)
				*urls = append(*urls, resolvedURL.String())
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		htmlParseHelper(c, baseUrl, urls)
	}
}
