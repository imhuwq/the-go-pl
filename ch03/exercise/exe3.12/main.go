package main

import (
	"flag"
	"fmt"
)

var s1 = flag.String("s1", "", "first string")
var s2 = flag.String("s2", "", "second string")

func isAnagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	l := len(s1)
	for o := 0; o <= l/2; o++ {
		if s1[o] != s2[l-o-1] {
			return false
		}
	}
	return true

}

func main() {
	flag.Parse()
	if isAnagrams(*s1, *s2) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
