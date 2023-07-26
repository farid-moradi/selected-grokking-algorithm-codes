package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sort"
)

var errDidntFindIt = errors.New("didn't find it")

func main() {
	// Ask for a number to search for
	fmt.Println("Search the list for ?")
	var n int
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		log.Fatal(err)

	}

	// making a slice of 100 random elements
	nList := make([]int, 100)
	for i := 0; i < 100; i++ {
		nList[i] = rand.Int() % 101
	}

	fmt.Println(nList)
	sort.Ints(nList)
	fmt.Println(nList)

	item, try, err := binarySearch(nList, n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("After %d tries\n", try)
	fmt.Println(item)
}

func binarySearch(list []int, n int) (item int, try int, err error) {
	ListLen := len(list)
	var mid, low, high int
	low = 0
	high = ListLen - 1

	if !sort.IntsAreSorted(list) {
		err = errors.New("the list is not sorted")
		item = -1
		try = -1
		return
	}

	for low <= high {
		// fmt.Println(mid, low, high)
		// time.Sleep(1 * time.Second)
		mid = (low + high) / 2
		try++
		if list[mid] == n {
			item = mid
			err = nil
			return
		} else if list[mid] < n {
			low = mid + 1
		} else if list[mid] > n {
			high = mid - 1
		} else {
			err = errors.New("internal error")
			return
		}
		// look how stupid was your first solution!
		// i = low + (high-low)/2 + 1
	}

	return -1, try, errDidntFindIt
}
