package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var s = flag.String("i", "", "input")
var l = flag.String("l", "256", "hash length, 256, 384 or 512, default 256")

func main() {
	flag.Parse()
	switch {
	case *l == "256":
		h := sha256.Sum256([]byte(*s))
		fmt.Printf("%T:\t%x\n", h, h)
	case *l == "384":
		h := sha512.Sum384([]byte(*s))
		fmt.Printf("%T:\t%x\n", h, h)
	case *l == "512":
		h := sha512.Sum512([]byte(*s))
		fmt.Printf("%T:\t%x\n", h, h)
	default:
		_, _ = fmt.Fprintf(os.Stderr, "exe4.2: err: invalid argument -l %v\n", *l)
		os.Exit(1)
	}
}
