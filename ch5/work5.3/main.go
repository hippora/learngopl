package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "work5.2: %v\n", err)
		os.Exit(1)
	}
	pText(doc)

}

func pText(n *html.Node) {
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}
	if n.Type == html.TextNode && len(strings.TrimSpace(n.Data)) != 0 {
		fmt.Println(strings.TrimSpace(n.Data))
	}
	if n.NextSibling != nil {
		pText(n.NextSibling)
	}
	if n.FirstChild != nil {
		pText(n.FirstChild)
	}
}
