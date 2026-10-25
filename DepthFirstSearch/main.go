package main

import "fmt"

var graph = [][]int{}
var visited = []bool{}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	graph = make([][]int, n)
	visited = make([]bool, n)

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		graph[a-1] = append(graph[a-1], b-1)
		graph[b-1] = append(graph[b-1], a-1)
	}

	dfs(0)

	for _, v := range visited {
		if !v {
			fmt.Println("The graph is not connected.")
			return
		}
	}

	fmt.Println("The graph is connected.")
}

func dfs(v int) {
	visited[v] = true

	for _, next := range graph[v] {
		if !visited[next] {
			dfs(next)
		}
	}
}
