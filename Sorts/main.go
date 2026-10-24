package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 3, 8, 1, 2}

	bubbleSort(nums)
	fmt.Printf("bubbleSort: %v\n", nums)

	selectionSort(nums)
	fmt.Printf("selectionSort: %v\n", nums)

	insertionSort(nums)
	fmt.Printf("insertionSort: %v\n", nums)

	mergeSort(nums)
	fmt.Printf("mergeSort: %v\n", nums)
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
