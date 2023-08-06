package main

import (
	"fmt"
	"math"
)

// t is the typed word, and s is the one that we think it might be the right word
func longestCommonSubsequence(s, t string) int {
	grid := make([][]int, len(s))
	var longest int
	for i, sv := range s {
		grid[i] = make([]int, len(t))
		for j, tv := range t {
			if i == 0 || j == 0 {
				if sv == tv {
					grid[i][j] = 1
				} else {
					if i == 0 {
						grid[i][j] = int(math.Max(float64(grid[i][j]), float64(grid[i][j-1])))
					} else {
						grid[i][j] = int(math.Max(float64(grid[i-1][j]), float64(grid[i][j])))
					}
				}
			} else {
				if sv == tv {
					grid[i][j] = grid[i-1][j-1] + 1
				} else {
					grid[i][j] = int(math.Max(float64(grid[i-1][j]), float64(grid[i][j-1])))
				}
			}
			if grid[i][j] > longest {
				longest = grid[i][j]
			}
		}
		fmt.Println(grid)
	}
	return longest
}
