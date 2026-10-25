package main

import "fmt"

type time struct {
	start int
	end   int
}

func main() {
	var n int
	fmt.Scan(&n)

	movies := make([]time, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&movies[i].start, &movies[i].end)
	}

	fmt.Println(intervalScheduling(movies))
}

func intervalScheduling(movies []time) int {
	if len(movies) == 0 {
		return 0
	}

	mergeSort(movies)

	i, count, lastEnd := 0, 0, 0
	for movies[i].start >= lastEnd {
		lastEnd = movies[i].end
		count++
		i++
	}

	return count
}

func mergeSort(movies []time) {
	if len(movies) <= 1 {
		return
	}

	mid := len(movies) / 2
	left := make([]time, mid)
	right := make([]time, len(movies)-mid)

	copy(left, movies[:mid])
	copy(right, movies[mid:])

	mergeSort(left)
	mergeSort(right)

	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i].end < right[j].end {
			movies[k] = left[i]
			i++
		} else {
			movies[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		movies[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		movies[k] = right[j]
		j++
		k++
	}
}
