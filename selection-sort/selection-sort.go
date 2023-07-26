package main

import (
	"fmt"
)

func main() {
	a := [7]int{1, 4, 17, 3, 2, 25, 9}
	fmt.Printf("%T\n", a)
	fmt.Println(a)
	selectionSort(a[:])
	fmt.Println(a)
}

func selectionSort(a []int) {
	var index int
	for i := 0; i < len(a)-1; i++ {
		// index represents the index of the smallest element
		index = i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[index] {
				index = j
			}
		}

		a[i], a[index] = a[index], a[i]

		// Report
		// fmt.Printf("\n")
		// fmt.Println(strings.Repeat("*", 12))
		// fmt.Println("founded in:", index)
		// fmt.Printf("step %d: %v\n", i+1, a)
		// fmt.Println(strings.Repeat("*", 12))
	}
}
