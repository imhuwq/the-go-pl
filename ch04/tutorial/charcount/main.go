package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "charcount: err: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\t\t\t\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t\t\t\t%d\n", c, n)
	}
	fmt.Printf("\nlen\t\t\t\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t\t\t\t%d\n", i, n)
		}
	}
	fmt.Printf("\ninvalid utf8 characters: %d\n", invalid)
}
