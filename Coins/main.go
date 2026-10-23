package main

import (
	"fmt"
	"os"
)

func main() {
	var A, B, C, X int
	fmt.Fscan(os.Stdin, &A, &B, &C, &X)

	count := 0

	for i := 0; i <= A; i++ {
		for j := 0; j <= B; j++ {
			for k := 0; k <= C; k++ {
				current := i*500 + j*100 + k*50
				if current == X {
					count++
				}
			}
		}
	}

	fmt.Printf("%d\n", count)
}
