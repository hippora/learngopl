package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%b\n%b\n", c1, c2)
	fmt.Printf("diff count : %d\n", diffCount(c1, c2))
}

func diffCount(c1, c2 [32]byte) int {
	var count int
	for i, v := range c1 {
		t1 := fmt.Sprintf("%08b", v)
		t2 := fmt.Sprintf("%08b", c2[i])
		fmt.Println(t1, "-", t2)
		for j := 0; j < len(t1); j++ {
			if t1[j] != t2[j] {
				count++
			}
		}
	}
	return count
}
