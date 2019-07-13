package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	strings := []string{"hello", "", "world", "go", "", "go", "", "go"}

	strings = nonempty(strings)
	fmt.Println(strings, len(strings), cap(strings))

	strings = nonempty2(strings)
	fmt.Println(strings, len(strings), cap(strings))
}
