package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	begin := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	end := time.Now()

	beginJoin := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	endJoin := time.Now()

	fmt.Println(end.Sub(begin), endJoin.Sub(beginJoin))
}
