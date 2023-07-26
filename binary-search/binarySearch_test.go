package main

import (
	"testing"
)

func TestBinarySearchIntNotExist(t *testing.T) {
	list := []int{1, 2, 3}
	_, _, err := binarySearch(list, 4)

	if err == nil {
		t.Fatal("test - shouldn't find the number")
	}
}

func TestBinarySearchIntExist(t *testing.T) {
	tests := []struct {
		list         []int
		expectedItem int
		item         int
		expectedErr  error
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2, nil},
		{[]int{1, 2, 3, 4}, 3, 2, nil},
		{[]int{1, 2, 3, 4, 5}, 5, 4, nil},
		{[]int{1, 2, 3, 4, 5}, 6, -1, errDidntFindIt},
	}

	for i, tt := range tests {
		item, _, err := binarySearch(tt.list, tt.expectedItem)
		if item != tt.item {
			t.Fatalf("error [%d] - item %d expected, getting %d\n", i, tt.item, item)
		}

		if err != tt.expectedErr {
			t.Fatalf("error [%d] - err %v expected, getting %v\n", i, tt.expectedErr, err)
		}
	}
}
