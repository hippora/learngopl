package main

import "fmt"

func main() {
	s := []byte("Hello,我的世  界!")
	reverse(&s)
	fmt.Println(string(s))
}

func reverse(b *[]byte) {
	r := []rune(string(*b))
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	*b = []byte(string(r))
}
