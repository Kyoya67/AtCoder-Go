package main

import "fmt"

type Queue []string

func (q *Queue) Enqueue(value string) {
	*q = append(*q, value)
}

func (q *Queue) Dequeue() {
	if empty := q.isEmpty(); empty {
		return
	}
	*q = (*q)[1:]
}

func (q *Queue) First() string {
	if empty := q.isEmpty(); empty {
		return ""
	}
	return (*q)[0]
}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}

func main() {
	var q Queue
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var order int
		fmt.Scan(&order)

		switch order {
		case 1:
			var title string
			fmt.Scan(&title)
			q.Enqueue(title)
		case 2:
			fmt.Println(q.First())
		case 3:
			q.Dequeue()
		}
	}
}
