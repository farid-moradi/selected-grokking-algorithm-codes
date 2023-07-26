package main

import "fmt"

func main() {
	a := []int{2, 12, 5, 3, 8, 1}
	quicksort(a)
	fmt.Println(a)
}

func quicksort(a []int) {
	if len(a) < 2 {
		return
	}

	pivot := 0
	// index is the index of the last bigger element before pivot
	index := 0
	for i := pivot + 1; i < len(a); i++ {
		if a[i] <= a[pivot] {
			index++
			a[index], a[i] = a[i], a[index]
		}
	}
	a[pivot], a[index] = a[index], a[pivot]
	pivot = index

	quicksort(a[:pivot])
	quicksort(a[pivot+1:])

}
