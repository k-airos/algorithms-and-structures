package main

import "fmt"

//Вывести все подмножества массива
func main() {
	var arr = []int{1, 2, 3}

	n := len(arr)

	for mask := 0; mask < (1 << n); mask++ { //1 * 2^n
		fmt.Print("{")
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				fmt.Print(arr[i], " ")
			}
		}
		fmt.Println("}")
	}

}

//mask = 000...0000 - empty
//mask = 000...0001 - {1}
//mask = 000...0010 - {2}
//mask = 000...0011 - {1 2}
//mask = 000...0100 - {3}
//mask = 000...0101 - {1,3}
//mask = 000...0110 - {2 3}
//mask = 000...0111 - {1 2 3}
