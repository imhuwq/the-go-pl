package main

import "fmt"

func unique(strings []string) []string {
	skipping := false
	out := strings[:0]
	for i, _ := range strings {
		if i == 0 {
			continue
		}

		if strings[i] != strings[i-1] {
			if skipping {
				skipping = false
			} else {
				out = append(out, strings[i-1])
			}
		} else {
			skipping = true
		}
	}

	if !skipping {
		out = append(out, strings[len(strings)-1])
	}

	return out
}

func main() {
	strings := []string{"hi", "hi", "hi", "hello", "world", "world", "you"}
	strings = unique(strings)
	fmt.Println(strings)
}
