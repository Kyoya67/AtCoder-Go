package main

import (
	"fmt"
)

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	fmt.Println(n, x)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Println(binarySearch(a, x))
}

func binarySearch(a []int, x int) int {
	mid := len(a) / 2

	if x == a[mid] {
		return mid
	} else if x < a[mid] {
		binarySearch(a[:mid], x)
	} else {
		binarySearch(a[mid:], x)
	}

	return -1
}
