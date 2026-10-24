package main

import (
	"fmt"
)

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Println(binarySearch(a, x, 0, len(a)-1) + 1)
}

func binarySearch(a []int, x int, left int, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2

	if x == a[mid] {
		return mid
	}

	if x < a[mid] {
		return binarySearch(a, x, left, mid-1)
	} else {
		return binarySearch(a, x, mid+1, right)
	}
}
