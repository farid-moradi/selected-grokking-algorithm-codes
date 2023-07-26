package main

import (
	"math/rand"
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
