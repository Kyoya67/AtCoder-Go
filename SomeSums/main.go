package main

import (
	"fmt"
	"os"
)

func main() {
	var N, A, B int
	fmt.Fscan(os.Stdin, &N, &A, &B)

	ans := 0

	for i := 1; i <= N; i++ {
		x := i
		digitSum := 0

		for x > 0 {
			digitSum += x % 10
			x /= 10
		}

		if A <= digitSum && digitSum <= B {
			ans += i
		}
	}

	fmt.Printf("%d\n", ans)
}
