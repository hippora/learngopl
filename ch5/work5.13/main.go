package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/hippora/learngo/ch5/links"
)

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	filtered := filterList(url, list)
	for _, v := range filtered {
		save(v)
	}

	return filtered
}

func filterList(orig string, list []string) []string {
	ourl, err := url.Parse(orig)
	if err != nil {
		log.Println(err)
		return nil
	}
	var filtered []string
	for _, v := range list {
		turl, err := url.Parse(v)
		if err != nil {
			log.Println(err)
			continue
		}
		// fmt.Printf("ourl:%s,turl:%s\n", ourl.Host, turl.Host)
		if ourl.Host == turl.Host {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}
func save(file string) {
	resp, err := http.Get(file)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	doc, err := url.Parse(file)
	if err != nil {
		log.Println(err)
		return
	}
	dir := filepath.Join("./", doc.Host, filepath.Clean(doc.Path))
	//for dividing dir path and file name
	base := "content"

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile(filepath.Join(dir, base), bs, os.ModePerm)

	if err != nil {
		log.Println(err)
		return
	}
}
