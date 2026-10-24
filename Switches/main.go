package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	s := make([][]int, m)
	p := make([]int, m)

	for i := 0; i < m; i++ {
		var k int
		fmt.Scan(&k)

		s[i] = make([]int, k)

		for j := 0; j < k; j++ {
			fmt.Scan(&s[i][j])
		}
	}

	for i := 0; i < m; i++ {
		fmt.Scan(&p[i])
	}

	ans := 0

	for bit := 0; bit < (1 << n); bit++ {
		ok := true

		for i := 0; i < m; i++ {
			count := 0

			for _, sw := range s[i] {
				if bit&(1<<(sw-1)) != 0 {
					count++
				}
			}

			if count%2 != p[i] {
				ok = false
				break
			}
		}

		if ok {
			ans++
		}
	}

	fmt.Println(ans)
}
