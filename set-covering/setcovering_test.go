package main

import "testing"

var statesNeeded []string
var stations map[string][]string

func init() {
	statesNeeded = []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}
	stations = make(map[string][]string)
	stations["kone"] = []string{"id", "nv", "ut"}
	stations["ktwo"] = []string{"wa", "id", "mt"}
	stations["kthree"] = []string{"or", "nv", "ca"}
	stations["kfour"] = []string{"nv", "ut"}
	stations["kfive"] = []string{"az", "ca"}
}

func BenchmarkSetcoveringExact(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findStationsGreedy(statesNeeded, stations)
		findStationsExact(statesNeeded, stations)
	}
}

func BenchmarkSetcoveringGreedy(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findStationsGreedy(statesNeeded, stations)
	}
}
