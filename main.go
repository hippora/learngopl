package main

import (
	//. "fmt"
	"fmt"
	"math"
	"unicode/utf8"
)

func main() {
	fmt.Printf("%d,%#[1]o,%#[1]x,%[1]b\n", 1234)
	fmt.Printf("%c,%[1]q\n", 'a')
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!
	fmt.Printf("%#3.8f,%[1]e,%[1]g\n", 3.141592653)
	fmt.Println(math.Sqrt(16.16))
	s := "Hello, 世界"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	r, size := utf8.DecodeRuneInString("世界")
	fmt.Printf("%c,%d\n", r, size)
	a, b := utf8.DecodeRune([]byte("世界"))
	fmt.Printf("%c,%d\n", a, b)
	for _, v := range []rune("プログラム") {
		fmt.Printf("%c\n", v)
	}

}
