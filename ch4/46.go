package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := []byte("Hello,我的  世界")
	s = repSpace(s)
	fmt.Println(string(s))

}

func repSpace(s []byte) []byte{
	l := len(s)
	for i:=0;i<l-1;{
		r,size := utf8.DecodeRune(s[i:])
		fmt.Printf("i=%d,r=%c,size=%d\n,s=%v\n",i,r,size,s)
		if unicode.IsSpace(r) {
			r1,size1 := utf8.DecodeRune(s[i+size:])
			fmt.Printf("r1=%c,size1=%d\n",r1,size1)
			if unicode.IsSpace(r1) {
				a := []byte("\n")
				copy(s[:i],append(a,s[i+size+size1:]...))
				i +=len(a)
				l = l-len(a)
				continue
			}
		}
		i +=size
	}
	return s[:l]
}
