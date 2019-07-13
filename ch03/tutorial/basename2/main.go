package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.String("n", "", "input name")

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	flag.Parse()
	fmt.Println(basename(*n))
}
