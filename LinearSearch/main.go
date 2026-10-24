package main

import (
	"fmt"
)

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	if linearSearch(n, x) {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
}

func linearSearch(n int, x int) bool {
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		if a == x {
			return true
		}
	}
	return false
}
