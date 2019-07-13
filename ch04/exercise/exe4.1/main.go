package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
)

var s1 = flag.String("s1", "", "first input")
var s2 = flag.String("s2", "", "second input")

func popCount(sha [32]byte) int {
	var val int = 0
	for _, v := range sha {
		for v != 0 {
			v &= v-1
			val++
		}
	}
	return val
}

func main () {
	flag.Parse()
	h1 := sha256.Sum256([]byte(*s1))
	h2 := sha256.Sum256([]byte(*s2))
	fmt.Printf("s1: %x\ns2: %x\n", h1, h2)
	fmt.Printf("bit difference: %d\n", popCount(h1) - popCount(h2))
}
