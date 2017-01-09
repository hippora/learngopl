package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
	for i,v := range os.Args {
		fmt.Printf("index:%d,arg:%s\n", i, v)
	}
}
