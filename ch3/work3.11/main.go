package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("123"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567.00001"))
	fmt.Println(comma("123456.00002"))
	fmt.Println(comma("+1234567"))
	fmt.Println(comma("-123456"))
	fmt.Println(comma("+1234567.00001"))
	fmt.Println(comma("-123456.00002"))
}

func comma(sf string) string {
	dot := strings.Index(sf, ".")
	s := sf
	if dot > 0 {
		s = s[:dot]
	}
	var buf bytes.Buffer
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		s = s[1:]
	}
	l := len(s)
	for i := 0; i < l; i++ {
		if (i-l%3)%3 == 0 && i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	if dot > 0 {
		buf.WriteString(sf[dot:])
	}
	return buf.String()
}
