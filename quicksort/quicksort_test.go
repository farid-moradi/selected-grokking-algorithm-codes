package main

import (
	"log"
	"math/rand"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	list := make([]int, 100)
	for i := range list {
		list[i] = rand.Int() % 100
	}
	listCheck := make([]int, 100)
	copy(listCheck, list)
	sort.Ints(listCheck)
	quicksort(list)

	for i, v := range list {
		if v != listCheck[i] {
			log.Fatalln("expected these 2 to be the same", list, listCheck)
		}
	}
}

func TestQuickSortRandomPivot(t *testing.T) {
	list := make([]int, 100)
	for i := range list {
		list[i] = rand.Int() % 100
	}
	listCheck := make([]int, 100)
	copy(listCheck, list)
	sort.Ints(listCheck)
	quickSortRandomPivot(list)

	for i, v := range list {
		if v != listCheck[i] {
			log.Fatalln("expected these 2 to be the same", list, listCheck)
		}
	}
}
