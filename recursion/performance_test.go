package main

import "testing"

func BenchmarkRecursion(b *testing.B) {
	countdown(12)
}

func BenchmarkSimple(b *testing.B) {
	countdownSimple(12)
}

func BenchmarkFact(b *testing.B) {
	fact(150)
}

func BenchmarkFactSimple(b *testing.B) {
	factSimple(150)
}
