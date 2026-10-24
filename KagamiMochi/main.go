package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Fscan(os.Stdin, &n)

	seen := make(map[int]bool)

	for i := 0; i < n; i++ {
		var f int
		fmt.Fscan(os.Stdin, &f)

		if seen[f] == true {
			continue
		}

		seen[f] = true
	}

	fmt.Printf("%d\n", len(seen))

	// nums := []int{}

	// for i := 0; i < n; i++ {
	// 	var f int
	// 	fmt.Fscan(os.Stdin, &f)
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
