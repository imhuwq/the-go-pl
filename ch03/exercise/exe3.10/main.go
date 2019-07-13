package main

import (
	"bytes"
	"flag"
	"fmt"
)

var n = flag.String("n", "", "input string number")

func comma(s string) string {
	var buf bytes.Buffer

	pos := len(s) % 3
	if pos == 0 {
		pos = 3
	}
	buf.WriteString(s[:pos])
	for i := pos; i < len(s); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[pos : pos+3])
	}
	return buf.String()
}

func main() {
	flag.Parse()
	fmt.Println(comma(*n))
}
