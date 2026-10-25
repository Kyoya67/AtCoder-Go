package main

import "fmt"

var graph = [][]int{}
var dist = []int{}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	graph = make([][]int, n)
	dist = make([]int, n)

	for i := range dist {
		dist[i] = -1
	}

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		graph[a-1] = append(graph[a-1], b-1)
		graph[b-1] = append(graph[b-1], a-1)
	}

	bfs(0)

	for i, d := range dist {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(d)
	}
	fmt.Println()
}

func bfs(v int) {
	queue := []int{v}
	dist[v] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, next := range graph[current] {
			if dist[next] == -1 {
				dist[next] = dist[current] + 1
				queue = append(queue, next)
			}
		}
	}
}

// 1 4
// 2 3
// 3 4
// 5 6
// 1 2
// 2 4

// 1:4|,2|
// 2:3|,1|,4|
// 3:2|,4|
// 4:1|,3|,2|

// 5:6|
// 6:5|
