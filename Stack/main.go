package main

import "fmt"

type Stack []string

func (s *Stack) Push(value string) {
	*s = append(*s, value)
}

func (s *Stack) Pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *Stack) Top() string {
	return (*s)[len(*s)-1]
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func main() {
	var s Stack
	var n int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var order int
		fmt.Scan(&order)

		switch order {
		case 1:
			var title string
			fmt.Scan(&title)
			s.Push(title)
		case 2:
			if !s.isEmpty() {
				fmt.Println(s.Top())
			}
		case 3:
			if !s.isEmpty() {
				s.Pop()
			}
		}

	}
}
