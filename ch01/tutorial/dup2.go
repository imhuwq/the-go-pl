package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLinesFromFile(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func countLinesFromFiles(files []string, counts map[string]int) {
	for _, path := range files {
		file, err := os.Open(path)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLinesFromFile(file, counts)
	}
}

func printDupLines(counts map[string]int) {
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLinesFromFile(os.Stdin, counts)
	} else {
		countLinesFromFiles(files, counts)
	}

	printDupLines(counts)
}
