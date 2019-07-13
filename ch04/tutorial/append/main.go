package main

import "fmt"

type Month int

var Arr = [...]int{
	1:  1,
	2:  2,
	3:  3,
	4:  4,
	5:  5,
	6:  6,
	7:  7,
	8:  8,
	9:  9,
	10: 10,
	11: 11,
	12: 12}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func printIntSliceInfo(arr []int) {
	fmt.Printf("type: %T\tlen: %d\t\tcap: %d\t\tdata: %v\n", arr, len(arr), cap(arr), arr)
}

func main() {
	arr := Arr[:]
	printIntSliceInfo(arr)
	arr = appendInt(arr, 99)
	printIntSliceInfo(arr)

	arr2to5 := Arr[2:6]
	printIntSliceInfo(arr2to5)
	arr2to5 = appendInt(arr2to5, 999)
	printIntSliceInfo(arr2to5)

	arr3to6 := Arr[3:7]
	printIntSliceInfo(arr3to6)  // watch out the element 999
}
