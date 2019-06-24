package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := []byte("Hello,我        的  世 界!")
	fmt.Printf("%s,len=%d,cap=%d\n", string(s), len(s), cap(s))
	s = repSpace(s)
	fmt.Printf("%s,len=%d,cap=%d\n", string(s), len(s), cap(s))
}

func repSpace(s []byte) []byte {
	l := len(s)
	var pr rune
	for i := 0; i < l; {
		r, size := utf8.DecodeRune(s[i:])
		fmt.Printf("i=%d,r=%c,size=%d,s=%v\n", i, r, size, s)
		if unicode.IsSpace(r) && unicode.IsSpace(pr) {
			copy(s[i:], s[i+size:])
			l -= size
			continue
		}
		pr = r
		i += size
	}
	return s[:l]
}
