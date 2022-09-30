package util

import (
	"log"
	"strings"

	"golang.org/x/net/html"
)

func FindAnchors(body string) []string {
	parsedbody, err := html.Parse(strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	var anchor func(*html.Node)

	var anchorSlice []string

	anchor = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					anchorSlice = append(anchorSlice, a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			anchor(c)
		}
	}
	anchor(parsedbody)

	return anchorSlice
}
