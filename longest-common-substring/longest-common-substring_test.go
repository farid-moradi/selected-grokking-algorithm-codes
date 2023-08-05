package main

import (
	"log"
	"testing"
)

func TestLongestCommonString(t *testing.T) {
	tests := []struct {
		s, t string
		l    int
	}{
		{"fish", "hish", 3},
		{"fish", "vista", 2},
	}

	for i, t := range tests {
		if t.l != longestCommonSubstring(t.s, t.t) {
			log.Fatalf("test (%d): expected (%d) for (%s) and (%s): got (%d)\n", i, t.l, t.s, t.t, longestCommonSubstring(t.s, t.t))
		}
	}
}
