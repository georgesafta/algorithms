package search_test

import (
	"testing"

	mysearch "github.com/georgesafta/algorithms/search"
)

func TestLinearSearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 5, 7, 8, 9}

	index := mysearch.LinearSearch(3, arr)
	if index != 2 {
		t.Fatalf("Response not matching, expected : %v, got : %v", 2, index)
	}

	index = mysearch.LinearSearch(6, arr)
	if index != -1 {
		t.Fatalf("Response not matching, expected : %v, got : %v", -1, index)
	}
}
