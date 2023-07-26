package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nList := make([]int, 100)
	for i := 0; i < 100; i++ {
		nList[i] = rand.Int() % 101
	}
	fmt.Println(nList)
	fmt.Println("sum:", sum(nList))
	fmt.Println("count:", count(nList))
	fmt.Println("max:", findMax(nList))
}

func count(a []int) int {
	if len(a) == 0 {
		return 0
	} else {
		return 1 + count(a[1:])
	}
}

func findMax(a []int) int {
	if len(a) == 1 {
		return a[0]
	} else {
		if a[0] > a[1] {
			a[0], a[1] = a[1], a[0]
		}
		return findMax(a[1:])
	}
}

func sum(a []int) int {
	if len(a) == 0 {
		return 0
	} else {
		return a[0] + sum(a[1:])
	}
}
