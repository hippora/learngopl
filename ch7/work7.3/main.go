package main

import (
	"fmt"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	var s []string
	if t.left != nil {
		s = append(s, t.left.String())
	}
	s = append(s, fmt.Sprint(t.value))
	if t.right != nil {
		s = append(s, t.right.String())
	}
	return strings.Join(s, "->")
}

func main() {
	t := new(tree)
	add(t, 99)
	add(t, 1)
	add(t, 55)
	add(t, 3)
	add(t, 44)
	add(t, 2)

	s := appendValues([]int{88}, t)
	// Sort(s)
	fmt.Println(s)
	fmt.Println(t)

}
