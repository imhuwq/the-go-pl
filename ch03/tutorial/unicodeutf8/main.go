package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}
