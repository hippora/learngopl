package main

import "fmt"

func main() {
	fmt.Println(max(1, 2, 3, 4, 5, 99, 88))
	fmt.Println(max1(1, 7, 2, 5, 3, 4))
}

func max(vals ...int) int {
	if len(vals) == 0 {
		panic("no param!")
	}
	var result int
	for i, v := range vals {
		if i == 0 {
			result = v
			continue
		}
		if v > result {
			result = v
		}

	}
	return result
}
func max1(i int, vals ...int) int {
	result := i
	for _, v := range vals {
		if v > result {
			result = v
		}

	}
	return result
}
