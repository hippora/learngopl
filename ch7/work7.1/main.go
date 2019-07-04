package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// WordCounter Word Counter
type WordCounter int

// LineCounter Line Counter
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewReader(p))
	s.Split(bufio.ScanWords)
	for s.Scan() {
		*c++
	}
	return len(p), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewReader(p))
	s.Split(bufio.ScanLines)
	for s.Scan() {
		*c++
	}
	return len(p), nil
}

func main() {
	var c WordCounter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s. I'm John.", name)
	fmt.Println(c)
	var c2 LineCounter
	fmt.Fprintf(&c2, "hello,\n %s.\nI'm John.\n", name)
	fmt.Println(c2)

}
