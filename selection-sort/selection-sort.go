package main

import "fmt"

func main() {
	a := [7]int{1, 4, 17, 3, 2, 25, 9}
	fmt.Printf("%T\n", a)

	selectionSort(a[:])
	fmt.Println(a)
}

func selectionSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[j] > a[i] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
