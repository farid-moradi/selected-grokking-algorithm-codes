package main

import (
	"log"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		s, t string
		l    int
	}{
		{"fosh", "fort", 2},
		{"fosh", "fish", 3},
	}

	for i, t := range tests {
		if t.l != longestCommonSubsequence(t.s, t.t) {
			log.Fatalf("test (%d): expected (%d) for (%s) and (%s): got (%d)\n", i, t.l, t.s, t.t, longestCommonSubsequence(t.s, t.t))
		}
	}
}
