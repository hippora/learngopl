package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		wordfreq(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
				continue
			}
			wordfreq(f, counts)
			f.Close()
		}
	}
	for word, n := range counts {
		fmt.Printf("%d\t%s\n", n, word)
	}
}

func wordfreq(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}
}
