package main

import (
	"log"
	"math/rand"
	"sort"
	"testing"
)

func BenchmarkSelection100(b *testing.B) {
	// b.ReportAllocs()
	b.ResetTimer()
	nList := make([]int, 100)
	for i := 0; i < 100; i++ {
		nList[i] = rand.Int() % 101
	}
	selectionSort(nList)
}

func BenchmarkSelection1000(b *testing.B) {
	// b.ReportAllocs()
	b.ResetTimer()
	nList := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nList[i] = rand.Int() % 1001
	}
	selectionSort(nList)
}

func BenchmarkSelection10000(b *testing.B) {
	// b.ReportAllocs()
	b.ResetTimer()
	nList := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		nList[i] = rand.Int() % 10001
	}
	selectionSort(nList)
}

func TestSelectionSort(t *testing.T) {
	list := make([]int, 100)
	checkList := make([]int, 100)

	for i := 0; i < 100; i++ {
		list[i] = rand.Int() % 100
	}
	copy(checkList, list)

	selectionSort(list)
	sort.Ints(checkList)

	for i, v := range list {
		if v != checkList[i] {
			log.Fatalln("expected these 2 to be the same", list, checkList)
		}
	}
}
