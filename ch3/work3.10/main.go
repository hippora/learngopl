package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1234567890123"))
}

func comma(s string) string {
	l := len(s)
	if l <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i := 0; i < l; i++ {
		if (i-l%3)%3 == 0 && i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}
