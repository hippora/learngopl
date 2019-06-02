package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s, sep string
	//1
	sep = "-"
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}
	fmt.Println(s)
	secs := time.Since(start).Nanoseconds()
	fmt.Println("duration1:", secs)
	//2
	s = ""
	start = time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg
	}
	fmt.Println(s)
	secs = time.Since(start).Nanoseconds()
	fmt.Println("duration2:", secs)
	//3
	start = time.Now()
	s = strings.Join(os.Args[1:], "-")
	fmt.Println(s)
	secs = time.Since(start).Nanoseconds()
	fmt.Println("duration3:", secs)

}
