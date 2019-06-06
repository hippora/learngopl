package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hippora/learngo/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		switch {
		case time.Since(item.CreatedAt) < time.Duration(time.Hour*24*30):
			fmt.Printf("#1=%v %-5d %9.9s %.55s\n",
				item.CreatedAt, item.Number, item.User.Login, item.Title)
		case time.Since(item.CreatedAt) < time.Duration(time.Hour*24*365):
			fmt.Printf("#2=%v %-5d %9.9s %.55s\n",
				item.CreatedAt, item.Number, item.User.Login, item.Title)
		default:
			fmt.Printf("#3=%v %-5d %9.9s %.55s\n",
				item.CreatedAt, item.Number, item.User.Login, item.Title)
		}

	}
}
