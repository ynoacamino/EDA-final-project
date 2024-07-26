package algorithms

import "testing"

func TestQuickSortStart(t *testing.T) {
	arr := []int{10, 7, 8, 9, 1, 5}
	compareTo := func(a, b int) int {
		if a == b {
			return 0
		} else if a > b {
			return 1
		} else {
			return -1
		}
	}

	expected := []int{1, 5, 7, 8, 9, 10}
	result := QuickSortStart(arr, compareTo)

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %v, but got %v", expected[i], v)
		}
	}
}
