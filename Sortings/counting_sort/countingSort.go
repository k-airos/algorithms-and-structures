package main

import (
	"fmt"
	"math/rand"
)

func countingSort(arr *[]int) {
	minimum := (*arr)[0]
	maximum := (*arr)[0]
	for i := 0; i < len(*arr); i++ {
		if (*arr)[i] > maximum {
			maximum = (*arr)[i]
		}
		if (*arr)[i] < minimum {
			minimum = (*arr)[i]
		}
	}
	bucket := make([]int, maximum-minimum+1)
	for i := 0; i < len(*arr); i++ {
		bucket[(*arr)[i]-minimum]++
	}
	*arr = []int{}
	for i := 0; i < len(bucket); i++ {
		count := bucket[i] //i+minimum
		for j := 0; j < count; j++ {
			*arr = append(*arr, i+minimum)
		}
	}
}

func main() {
	var arr []int
	n := rand.Intn(100)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Printf("Old array: %v\n", arr)
	countingSort(&arr)
	fmt.Printf("New arr %v", arr)
}
