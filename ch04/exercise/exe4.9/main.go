package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		count[word]++
	}

	_, _ = fmt.Fprintf(os.Stdout, "word\tfreq\n")
	for w, f := range count {
		_, _ = fmt.Fprintf(os.Stdout, "%s\t%d\n", w, f)
	}
}
