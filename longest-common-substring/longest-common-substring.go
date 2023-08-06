package main

import "fmt"

// t is the typed word, and s is the one that we think it might be the right word
func longestCommonSubstring(s, t string) int {
	grid := make([][]int, len(s))
	var longest int
	for i, sv := range s {
		grid[i] = make([]int, len(t))
		for j, tv := range t {
			// if both characters are the same, the cell value becomes the length of the previous common substring just before now plus 1
			if sv == tv {
				if i > 0 && j > 0 {
					grid[i][j] = grid[i-1][j-1] + 1
				} else {
					grid[i][j] = 1
				}
				if grid[i][j] > longest {
					longest = grid[i][j]
				}
			}
		}
	}
	return longest
}
