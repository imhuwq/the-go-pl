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

func PopCountByShift(x uint64) int {
	var val uint64 = 0
	for i := 0; i < 64; i++ {
		val += x&1
		x = x>>1
	}
	return int(val)
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

		var countByLookUp int = 0
		var costByLookUp int64 = 0
		startLookUp := time.Now()
		for i := 0; i < runCount; i++ {
			startLookUp = time.Now()
			countByLookUp = PopCount(uint64(val))
			costByLookUp += time.Since(startLookUp).Nanoseconds()
		}
		_, _ = fmt.Fprintf(os.Stdout, "popcount: lookup out: %d, cost: %d\n", countByLookUp, costByLookUp/int64(runCount))

		var countByShift int = 0
		var costByShift int64 = 0
		startByShift := time.Now()
		for i:= 0; i < runCount; i++ {
			startByShift = time.Now()
			countByShift = PopCountByShift(uint64(val))
			costByShift += time.Since(startByShift).Nanoseconds()
		}
		_, _ = fmt.Fprintf(os.Stdout, "popcount:  shift out: %d, cost: %d\n", countByShift, costByShift/int64(runCount))

	}
}
