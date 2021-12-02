package main

import (
	"fmt"
	"math/rand"
)

func merge(left, right []int) []int {
	merged := make([]int, 0)
	var leftIndex, rightIndex int

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] <= right[rightIndex] {
			merged = append(merged, left[leftIndex])
			leftIndex++
		} else {
			merged = append(merged, right[rightIndex])
			rightIndex++
		}
	}
	for leftIndex < len(left) {
		merged = append(merged, left[leftIndex])
		leftIndex++
	}
	for rightIndex < len(right) {
		merged = append(merged, right[rightIndex])
		rightIndex++
	}
	return merged
}

func MergeSort(arr *[]int) {
	if len(*arr) <= 1 {
		return
	}
	left, right := make([]int, 0), make([]int, 0)
	n := len(*arr) / 2
	for i := 0; i < n; i++ {
		left = append(left, (*arr)[i])
	}
	for i := n; i < len(*arr); i++ {
		right = append(right, (*arr)[i])
	}

	MergeSort(&left)
	MergeSort(&right)

	*arr = merge(left, right)
}

func main() {
	var arr []int
	n := rand.Intn(100)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Printf("Old array: %v\n", arr)
	MergeSort(&arr)
	fmt.Printf("New arr %v", arr)
}
