package main

import (
	"fmt"
	"math/rand"
)

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	b := rand.Int()
	pivot := b % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}
func main() {
	var arr = []int{-4, 1, 7, 8, 6, 3, 1, 4}

	fmt.Printf("Old array: %v\n", arr)
	arr = quickSort(arr)
	fmt.Printf("New arr %v", arr)
}
