package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type xkcd struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	Safetitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

const filename = "xkcd.json"

func main() {
	if len(os.Args[1:]) == 0 {
		log.Fatal("no keywords,exit")
		os.Exit(1)
	}
	_, err := os.Stat(filename)
	if err != nil || os.IsNotExist(err) {
		downloadAll()
	}
	db := loadAll()
	for _, v := range db {
		for _, arg := range os.Args[1:] {
			if strings.Contains(v.Transcript, arg) {
				fmt.Printf("keyword: %s,Num:%3d,Title:%s, URL:=%s\n", arg, v.Num, v.Title, v.Img)
			}
		}
	}
}

func downloadAll() {
	var result xkcd

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("open file error:%s", err)
	}
	defer f.Close()
	for i := 1; i < 2150; i++ {
		url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("http get error:%v\n", err)
			resp.Body.Close()
			continue
		}
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("status %d,error:%v\n", resp.StatusCode, err)
			resp.Body.Close()
			continue
		}
		buf, err := ioutil.ReadAll(resp.Body)
		if err := json.Unmarshal(buf, &result); err != nil {
			fmt.Printf("json decode error:%v\n", err)
			resp.Body.Close()
			continue
		}
		f.Write(buf)
		resp.Body.Close()
	}
}

func loadAll() []xkcd {
	var db []xkcd
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("open file error:%s", err)
	}
	jd := json.NewDecoder(bytes.NewReader(data))
	var x xkcd
	for jd.More() {
		jd.Decode(&x)
		db = append(db, x)
	}
	return db
}
