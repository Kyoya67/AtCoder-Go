package main

import "fmt"

func main() {
	m := make(map[string]int)

	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var order int
		fmt.Scan(&order)

		switch order {
		case 1:
			var name string
			var score int

			fmt.Scan(&name, &score)
			m[name] = score
		case 2:
			var name string

			fmt.Scan(&name)
			fmt.Println(m[name])
		}
	}
}
