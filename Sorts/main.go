package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := []int{5, 3, 8, 1, 2}

	// bubbleSort(nums)
	// fmt.Printf("bubbleSort: %v\n", nums)

	// selectionSort(nums)
	// fmt.Printf("selectionSort: %v\n", nums)

	// insertionSort(nums)
	// fmt.Printf("insertionSort: %v\n", nums)

	mergeSort(nums)
	fmt.Printf("mergeSort: %v\n", nums)

	// quickSort(nums)
	// fmt.Printf("quickSort: %v\n", nums)

	// heapSort(nums)
	// fmt.Printf("heapSort: %v\n", nums)

	// bogosort(nums)
	// fmt.Printf("bogosort: %v\n", nums)
}

func bubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := n - 1; j > i; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}

	// for i := n - 1; i > 0; i-- {
	// 	for j := 0; j < i; j++ {
	// 		if nums[j] > nums[j+1] {
	// 			nums[j], nums[j+1] = nums[j+1], nums[j]
	// 		}
	// 	}
	// }
}

func selectionSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
}

func insertionSort(nums []int) {
	n := len(nums)
	for i := 1; i < n; i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
}

func mergeSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	mid := len(nums) / 2
	left := make([]int, mid)
	right := make([]int, len(nums)-mid)

	copy(left, nums[:mid])
	copy(right, nums[mid:])

	// left := nums[:mid]
	// right := nums[mid:]

	mergeSort(left)
	mergeSort(right)

	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			nums[k] = left[i]
			i++
		} else {
			nums[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		nums[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		nums[k] = right[j]
		j++
		k++
	}

}

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	pivot := nums[len(nums)/2]
	left := []int{}
	right := []int{}
	mid := []int{}
	for _, v := range nums {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			mid = append(mid, v)
		}
	}
	quickSort(left)
	quickSort(right)
	copy(nums, append(append(left, mid...), right...))
}

func heapSort(nums []int) {
	n := len(nums)
	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}
	// Extract elements from heap
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, i, 0)
	}
}

func heapify(nums []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && nums[left] > nums[largest] {
		largest = left
	}
	if right < n && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		heapify(nums, n, largest)
	}
}

func bogosort(nums []int) {
	for !isSorted(nums) {
		shuffle(nums)
	}
}

func isSorted(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			return false
		}
	}
	return true
}

func shuffle(nums []int) {
	for i := range nums {
		j := rand.Intn(i + 1)
		nums[i], nums[j] = nums[j], nums[i]
	}
}
