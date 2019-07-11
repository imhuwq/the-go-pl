package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
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

func PopCountByLoop(x uint64) int {
	val := 0
	for i := 0; i < 8; i++ {
		val += int(pc[byte(x>>(uint(i)*8))])
	}
	return val
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
		const runCount int = 10

		var countNoLoop int = 0
		var costNoLoop int64 = 0
		startNoLoop := time.Now()
		for i := 0; i < runCount; i++ {
			startNoLoop = time.Now()
			countNoLoop = PopCount(uint64(val))
			costNoLoop += time.Since(startNoLoop).Nanoseconds()
		}
		_, _ = fmt.Fprintf(os.Stdout, "popcount: no loop out: %d, cost: %d\n", countNoLoop, costNoLoop/int64(runCount))

		var countLoop int = 0
		var costLoop int64 = 0
		startLoop := time.Now()
		for i:= 0; i < runCount; i++ {
			startLoop = time.Now()
			countLoop = PopCountByLoop(uint64(val))
			costLoop += time.Since(startLoop).Nanoseconds()
		}
		_, _ = fmt.Fprintf(os.Stdout, "popcount:    loop out: %d, cost: %d\n", countLoop, costLoop/int64(runCount))

	}
}
