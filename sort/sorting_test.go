package sort_test

import (
	"testing"

	mysort "github.com/georgesafta/algorithms/sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (e ByAge) Len() int { return len(e) }
func (e ByAge) Compare(i, j int) int8 {
	if e[i].Age == e[j].Age {
		return 0
	}
	if e[i].Age > e[j].Age {
		return -1
	}

	return 1
}
func (e ByAge) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

func TestInsertionSort(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 5, 7, 8, 9}
	sorted_list := mysort.InsertionSort(3, 5, 2, 1, 5, 7, 8, 4, 9)

	for i := 0; i < len(expected); i++ {
		if sorted_list[i] != expected[i] {
			t.Fatalf("Response not matching, expected : %v, got : %v", expected[i], sorted_list[i])
		}
	}
}

func TestSelectionSort(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 5, 7, 8, 9}
	sorted_list := mysort.SelectionSort(3, 5, 2, 1, 5, 7, 8, 4, 9)

	for i := 0; i < len(expected); i++ {
		if sorted_list[i] != expected[i] {
			t.Fatalf("Response not matching, expected : %v, got : %v", expected[i], sorted_list[i])
		}
	}
}

func TestBubbleSort(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 5, 7, 8, 9}
	sorted_list := mysort.BubbleSort(3, 5, 2, 1, 5, 7, 8, 4, 9)

	for i := 0; i < len(expected); i++ {
		if sorted_list[i] != expected[i] {
			t.Fatalf("Response not matching, expected : %v, got : %v", expected[i], sorted_list[i])
		}
	}
}

func TestMergeSort(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 5, 7, 8, 9}
	s := []int{3, 5, 2, 1, 5, 7, 8, 4, 9}
	mysort.MergeSort(s...)

	for i := 0; i < len(expected); i++ {
		if s[i] != expected[i] {
			t.Fatalf("Response not matching, expected : %v, got : %v", expected[i], s[i])
		}
	}
}

func TestCustomInsertionSort(t *testing.T) {
	people := []Person{
		{"John", 31},
		{"Mary", 22},
		{"Paul", 25},
	}

	mysort.InsertionSortCustom(ByAge(people))
	for i := 1; i < len(people); i++ {
		if people[i-1].Age > people[i].Age {
			t.Fatalf("%v is before : %v", people[i-1].Age, people[i].Age)
		}
	}
}
