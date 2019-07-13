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
	letterCount := 0
	digitCount := 0
	invalid := 0
	var utflen [utf8.UTFMax + 1]int

	input := bufio.NewReader(os.Stdin)
	for {
		r, n, err := input.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "exe4.8: err: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letterCount++
		} else if unicode.IsDigit(r) {
			digitCount++
		}
		utflen[n]++;
	}

	_, _ = fmt.Fprintf(os.Stdout, "letter count: %d\t\tdigit count: %d\t\tinvalid count: %d\n", letterCount, digitCount, invalid)
	_, _ = fmt.Fprintf(os.Stdout, "length\tcount\n")
	for i, v := range utflen {
		if i > 0 {
			_, _ = fmt.Fprintf(os.Stdout, "%d\t%d\n", i, v)
		}
	}
}
