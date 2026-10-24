package main

import (
	"fmt"
	"os"
)

func main() {
	var n, x int
	fmt.Fscan(os.Stdin, &n, &x)

	if linearSearch(n, x) {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
}

func linearSearch(n int, x int) bool {
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(os.Stdin, &a)
		if a == x {
			return true
		}
	}
	return false
}
