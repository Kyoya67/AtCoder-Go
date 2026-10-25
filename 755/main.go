package main

import "fmt"

var n int
var count = 0

func main() {
	fmt.Scan(&n)

	dfs(0)

	fmt.Println(count)
}

func dfs(x int) {
	if x > n {
		return
	}

	if is753(x) {
		count++
	}

	dfs(x*10 + 3)
	dfs(x*10 + 5)
	dfs(x*10 + 7)
}

func is753(x int) bool {
	var hasDigit = map[int]bool{3: false, 5: false, 7: false}

	for x > 0 {
		switch x % 10 {
		case 3:
			hasDigit[3] = true
		case 5:
			hasDigit[5] = true
		case 7:
			hasDigit[7] = true
		default:
			return false
		}
		x /= 10
	}

	return hasDigit[3] && hasDigit[5] && hasDigit[7]
}
