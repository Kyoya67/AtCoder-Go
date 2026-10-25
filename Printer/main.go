package main

import "fmt"

func main() {
	var n int
	var k int64
	fmt.Scan(&n, &k)

	a := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Println(binarySearch(k, a))
}

func countPrint(t int64, a []int64) int64 {
	var count int64

	for _, v := range a {
		count += t / v
	}

	return count
}

func binarySearch(k int64, a []int64) int64 {
	left := int64(0)
	right := int64(1)

	for countPrint(right, a) < k {
		right *= 2
	}

	for right-left > 1 {
		mid := (left + right) / 2

		if countPrint(mid, a) >= k {
			right = mid
		} else {
			left = mid
		}
	}

	return right
}
