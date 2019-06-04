package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename("a/b/c.go")) // "c"
	fmt.Println(basename("c.d.go"))   // "c.d"
	fmt.Println(basename("abc"))      // "abc"
	s := "abcd"
	b := []byte(s)
	fmt.Println(b, string(b), strings.Fields("1 2 3 4 a"), strings.Split("1,2,3,4,5", ","))
	fmt.Println(bytes.Fields([]byte("1 2 3 4 5")))
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot > 0 {
		s = s[:dot]
	}
	return s
}
