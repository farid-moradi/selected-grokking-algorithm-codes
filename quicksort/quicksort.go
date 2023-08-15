package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := []int{2, 12, 5, 3, 8, 1}
	// quicksort(a)
	quickSortRandomPivot(a)
	fmt.Println(a)
}

func quickSortRandomPivot(a []int) {
	if len(a) < 2 {
		return
	}

	pivot := rand.Int() % len(a)
	index := 0

	for i := 0; i < len(a); i++ {
		if index > pivot && i <= len(a)-1 {
			a[pivot], a[pivot+1] = a[pivot+1], a[pivot]
			pivot++
		}
		if a[i] <= a[pivot] && i != pivot {
			if index == pivot {
				pivot = i
			}
			a[index], a[i] = a[i], a[index]
			index++
		}
	}
	a[index], a[pivot] = a[pivot], a[index]
	pivot = index

	quickSortRandomPivot(a[:pivot])
	quickSortRandomPivot(a[pivot+1:])
}

func quicksort(a []int) {
	if len(a) < 2 {
		return
	}

	pivot := 0
	// index of the biggest element before pivot
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
