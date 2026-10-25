package main

import "fmt"

type Books []string

func (b *Books) Push(value string) {
	*b = append(*b, value)
}

func (b *Books) Pop() {
	*b = (*b)[:len(*b)-1]
}

func (b *Books) Top() string {
	return (*b)[len(*b)-1]
}

func (b *Books) isEmpty() bool {
	return len(*b) == 0
}

func main() {
	var b Books
	var q int

	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		var order int
		fmt.Scan(&order)

		switch order {
		case 1:
			var title string
			fmt.Scan(&title)
			b.Push(title)
		case 2:
			if !b.isEmpty() {
				fmt.Println(b.Top())
			}
		case 3:
			if !b.isEmpty() {
				b.Pop()
			}
		}

	}
}
