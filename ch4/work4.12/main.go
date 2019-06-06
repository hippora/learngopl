package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	var result xkcd
	resp, err := http.Get("https://xkcd.com/571/info.0.json")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(err)
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	fmt.Println(result)
}

func dlall() {

}
