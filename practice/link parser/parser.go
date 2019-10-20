package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	var links []Link
	nodes := linkNodes(doc)
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}

	return links, nil
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var result []*html.Node

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result = append(result, linkNodes(c)...)
	}
	return result
}

func buildLink(n *html.Node) Link {
	var result Link

	for _, attr := range n.Attr {
		if attr.Key == "href" {
			result.Href = attr.Val
			result.Text = getText(n)
			break
		}
	}
	return result
}

func getText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type != html.ElementNode {
		return ""
	}

	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += getText(c) + " "
	}
	return strings.Join(strings.Fields(text), " ")
}
