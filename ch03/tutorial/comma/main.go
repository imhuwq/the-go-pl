package main

import (
	"flag"
	"fmt"
)

var n = flag.String("n", "", "input integer string")

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	flag.Parse()
	fmt.Println(comma(*n))
}