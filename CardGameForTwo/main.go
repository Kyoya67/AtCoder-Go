package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	sumA := 0
	sumB := 0

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			sumA += nums[i]
		} else {
			sumB += nums[i]
		}
	}

	fmt.Println(sumA - sumB)
}
