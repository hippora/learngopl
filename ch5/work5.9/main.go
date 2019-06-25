package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(expand("foobar", ex))
}

func expand(s string, f func(string) string) string {
	return strings.ReplaceAll(s, "foo", f("foo"))
}

func ex(s string) string {
	return strings.ReplaceAll(s, "foo", "bar")
}
