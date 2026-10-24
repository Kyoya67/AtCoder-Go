package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	fmt.Fscan(os.Stdin, &n)

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(os.Stdin, &nums[i])
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
