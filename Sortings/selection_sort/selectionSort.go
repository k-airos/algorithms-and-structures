package main

import (
	"fmt"
	"math/rand"
)

func SelectionSort(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		minInd := i
		for j := i + 1; j < len(*arr); j++ {
			if (*arr)[j] < (*arr)[minInd] {
				minInd = j
			}
		}
		c := (*arr)[i]
		(*arr)[i] = (*arr)[minInd]
		(*arr)[minInd] = c
	}
}

func main() {
	var arr []int
	n := rand.Intn(100)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Printf("Old array: %v\n", arr)
	SelectionSort(&arr)

	fmt.Printf("New arr %v", arr)
}
