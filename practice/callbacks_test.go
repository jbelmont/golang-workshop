package practice

import (
	"testing"
)

func TestFilter(t *testing.T) {
	actual := filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(n int) bool {
		return n%2 == 1
	})

	expected := []int{1, 3, 5, 7, 9}
	for idx, val := range actual {
		if val != expected[idx] {
			t.Error("only odd numbers should be here")
		}
	}
}

func TestMapper(t *testing.T) {
	actual := mapper([]int{0, 1, 2, 3, 4, 5}, func(n int) int {
		return n
	})
	expected := []int{0, 1, 2, 3, 4, 5}
	for i, v := range actual {
		if v != expected[i] {
			t.Error("should return 0 1 2 3 4 5")
		}
	}
}
