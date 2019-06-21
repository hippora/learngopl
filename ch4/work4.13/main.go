package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

type poster struct {
	Title    string
	Poster   string
	Response string
}

func main() {
	var result poster
	addr := fmt.Sprintf("http://www.omdbapi.com/?apikey=dab55e93&t=%s", url.QueryEscape(strings.Join(os.Args[1:], " ")))
	buf, err := getData(addr)
	if err != nil {
		fmt.Printf("http get error:%v\n", err)
		os.Exit(1)
	}
	if err := json.Unmarshal(buf, &result); err != nil {
		fmt.Printf("json decode error:%v\n", err)
		os.Exit(1)
	}
	if result.Response == "False" {
		fmt.Println("No Response!")
		os.Exit(1)
	}
	filename := fmt.Sprintf("%s%s", result.Title, path.Ext(result.Poster))
	fmt.Printf("Title:%s\nPoster:%s\nfilename:%s\n", result.Title, result.Poster, filename)

	buf, err = getData(result.Poster)
	if err != nil {
		fmt.Printf("http get error:%v\n", err)
		os.Exit(1)
	}
	ioutil.WriteFile(filename, buf, 0644)
}

func getData(url string) ([]byte, error) {
	// fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("http get error:%v\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("status %d,error:%v\n", resp.StatusCode, err)
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}
