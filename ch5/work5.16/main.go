package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(join("#", "aaa", "bbb", "ccc"))
}

func join(sep string, a ...string) string {
	return strings.Join(a, sep)
}
