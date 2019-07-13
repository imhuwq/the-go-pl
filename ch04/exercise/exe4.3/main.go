package main

import "fmt"

func reverse(arrPtr *[5]int) {
	for i, j := 0, len(*arrPtr)-1; i < j; i, j = i+1, j-1 {
		(*arrPtr)[i], (*arrPtr)[j] = (*arrPtr)[j], (*arrPtr)[i]
	}
}

func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	reverse(&arr)
	fmt.Println(arr)
}
