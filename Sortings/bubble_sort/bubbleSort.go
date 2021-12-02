package main

import (
	"fmt"
	"math/rand"
)

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		isSorted := true
		for j := 1; j < len(arr); j++ {
			if arr[j] < arr[j-1] {
				isSorted = false
				c := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = c
			}
		}
		if isSorted {
			return
		}
	}
}

func main() {
	n := rand.Intn(101)
	arr := make([]int, 0, n)
	fmt.Println(cap(arr))
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Printf("Old array: %v\n", arr)
	BubbleSort(arr)
	fmt.Printf("New array: %v\n", arr)
}
