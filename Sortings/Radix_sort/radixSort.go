package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func radixSort(arr *[]int) {
	buckets := make([][]int, 10)
	powerOfTen := 1
	maxLen := 0
	for i := 0; i < len(*arr); i++ {
		if len(strconv.Itoa(i)) > maxLen {
			maxLen = len(strconv.Itoa(i))
		}
	}
	for pow := 0; pow <= maxLen; pow++ {
		for _, elem := range *arr {
			buckets[elem/powerOfTen%10] = append(buckets[elem/powerOfTen%10], elem)
		}
		*arr = make([]int, 0)
		for i := 0; i < len(buckets); i++ {
			for j := 0; j < len(buckets[i]); j++ {
				*arr = append(*arr, buckets[i][j])
			}
			buckets[i] = nil
		}

		powerOfTen *= 10
	}

}

func main() {
	var arr []int
	n := rand.Intn(100)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Printf("Old array: %v\n", arr)
	radixSort(&arr)
	fmt.Printf("New arr %v", arr)
}
