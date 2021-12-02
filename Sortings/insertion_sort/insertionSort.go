package main

import (
	"fmt"
	"math/rand"
)

func InsertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		cur := arr[i]
		j := i - 1
		for j >= 0 && cur < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = cur
	}
}

func main() {
	var arr []int
	n := rand.Intn(100)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Printf("Old array: %v\n", arr)
	InsertionSort(arr)
	fmt.Printf("New arr: %v\n", arr)

}
