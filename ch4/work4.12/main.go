package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func main() {
	downloadAll()
}

func downloadAll() {
	var result xkcd
	for i := 1; i < 10; i++ {
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
		filename := fmt.Sprintf("ch4/work4.12/%d.json", i)
		ioutil.WriteFile(filename, buf, 0644)
		resp.Body.Close()
	}
}
