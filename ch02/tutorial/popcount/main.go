package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func intFromArgOrStdin() (int, bool) {
	if len(os.Args) >= 2 {
		val, err := strconv.Atoi(os.Args[1])
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "popcount: err parse argument: %v\n", err)
			return 0, false
		}
		return val, true
	} else {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		val, err := strconv.Atoi(input.Text())
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "popcount: err parse stdin: %v\n", err)
			return 0, false
		}
		return val, true
	}
}

func main() {
	if val, ok := intFromArgOrStdin(); ok {
		_, _ = fmt.Fprintf(os.Stdout, "popcount: out: %d\n", PopCount(uint64(val)))
	}
}
