package main

import (
	"bytes"
	"flag"
	"fmt"
	"strings"
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

func fractComma(s string) string {
	var buf bytes.Buffer
	i := 0
	for i = 0; i < len(s)-3; i += 3 {
		buf.WriteString(s[i : i+3])
		buf.WriteByte(',')
	}
	buf.WriteString(s[i:])
	return buf.String()
}

func enhancedComma(s string) string {
	prefix := ""
	suffix := ""
	if strings.HasPrefix(s, "+") {
		prefix = "+"
		s = s[1:]
	} else if strings.HasPrefix(s, "-") {
		prefix = "-"
		s = s[1:]
	}

	if pointIndex := strings.LastIndex(s, "."); pointIndex >= 0 {
		suffix = "." + fractComma(s[pointIndex+1:])
		s = s[:pointIndex]
	}

	return prefix + comma(s) + suffix
}

func main() {
	flag.Parse()
	fmt.Println(enhancedComma(*n))
}
