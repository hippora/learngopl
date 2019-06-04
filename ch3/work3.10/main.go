package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1234567890123"))
	fmt.Println(comma("123456789012"))
}

func comma(s string) string {
	l := len(s)
	var buf bytes.Buffer
	for i := 0; i < l; i++ {
		if (i-l%3)%3 == 0 && i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}
