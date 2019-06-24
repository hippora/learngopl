package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "work5.2: %v\n", err)
		os.Exit(1)
	}
	c := make(map[string]int, 0)
	countElement(c, doc)
	for i, v := range c {
		fmt.Printf("element: %s\t%d\n", i, v)
	}
}

func countElement(count map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	if n.NextSibling != nil {
		count = countElement(count, n.NextSibling)
	}
	if n.FirstChild != nil {
		count = countElement(count, n.FirstChild)
	}
	return count
}
