package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sum := 0
	var answer int

	for sec := 1; ; sec++ {
		for i := 0; i < n; i++ {
			if sec%a[i] == 0 {
				sum++
				if sum == k {
					answer = sec
					goto skip
				}
			}
		}
	}
skip:
	fmt.Println(answer)
}
