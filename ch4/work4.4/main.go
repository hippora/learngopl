package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	rotate(s)
	fmt.Println(s)
}

func rotate(s []int) {
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}
