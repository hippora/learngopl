package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "work5.7: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, v := range n.Attr {
			fmt.Printf(" %s=\"%s\" ", v.Key, v.Val)
		}
		if n.FirstChild == nil {
			fmt.Println("/>")
		} else {
			fmt.Println(">")
		}
		depth++
	} else if n.Type == html.TextNode {
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s <!-- %s -->\n", depth*2, "", n.Data)
	}
}
