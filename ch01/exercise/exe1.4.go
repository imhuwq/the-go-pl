package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLinesFromFile(f *os.File, counts map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		_, exists := counts[line]
		if !exists {
			counts[line] = make(map[string]bool)
		}
		counts[line][f.Name()] = true
	}
}

func countLinesFromFiles(files []string, counts map[string]map[string]bool) {
	for _, path := range files {
		file, err := os.Open(path)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLinesFromFile(file, counts)
	}
}

func printDupLines(counts map[string]map[string]bool) {
	for line, lineCounts := range counts {
		if len(lineCounts) > 1 {
			var paths []string
			for path, _ := range lineCounts {
				paths = append(paths, path)
			}
			fmt.Printf("%s\t%v\n", line, paths)
		}
	}
}

func main() {
	counts := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLinesFromFile(os.Stdin, counts)
	} else {
		countLinesFromFiles(files, counts)
	}

	printDupLines(counts)
}
