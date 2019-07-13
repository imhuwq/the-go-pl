package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%T\t%x\n%T\t%x\n%t\n", c1, c1, c2, c2, c1 == c2)
}
