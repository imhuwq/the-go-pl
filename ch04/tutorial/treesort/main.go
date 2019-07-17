package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func add(t *tree, v int) *tree {
	if t == nil {
		t = new(tree)
		t.value = v
		return t
	}

	if v > t.value {
		t.right = add(t.right, v)
	} else {
		t.left = add(t.left, v)
	}
	return t
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}

	appendValues(values[:0], root)
}


func main() {
	slice := []int {5, 8, 2, 6, 7, 1, 10, 62, 45, 1}
	Sort(slice)
	fmt.Println(slice)
}