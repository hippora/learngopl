package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("%s url id\n", os.Args[0])
	}
	url, id := os.Args[1], os.Args[2]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "work5.7: %v\n", err)
		os.Exit(1)
	}
	r := ElementByID(doc, id)
	if r == nil {
		fmt.Printf("cant find element by id:%s", id)
	} else {
		fmt.Println(r)
	}
}

// ElementByID return *html.Node
func ElementByID(doc *html.Node, id string) *html.Node {
	var r *html.Node
	f := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, v := range n.Attr {
				if v.Key == "id" && v.Val == id {
					r = n
					return true
				}
			}
		}
		return false
	}
	forEachNode(doc, "", f, nil)
	return r
}

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node) bool) {
	var b bool
	if pre != nil {
		b = pre(n)
	}
	for c := n.FirstChild; c != nil && !b; c = c.NextSibling {
		forEachNode(c, id, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int
