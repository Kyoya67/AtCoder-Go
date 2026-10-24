package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	seen := make(map[int]bool)

	for i := 0; i < n; i++ {
		var f int
		fmt.Scan(&f)

		if seen[f] == true {
			continue
		}

		seen[f] = true
	}

	fmt.Printf("%d\n", len(seen))

	// nums := []int{}

	// for i := 0; i < n; i++ {
	// 	var f int
	// 	fmt.Scan(&f)
	// 	for i := 0; i < len(nums); i++ {
	// 		if f == nums[i] {
	// 			goto skip
	// 		}
	// 	}
	// 	nums = append(nums, f)
	// skip:
	// }
	// fmt.Println(len(nums))
}
