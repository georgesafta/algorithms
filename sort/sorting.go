package sort

type Comparable interface {
	Compare(i, j int) int8
	Len() int
	Swap(i, j int)
}

// InsertionSortCustom sorts a type safisfying the Comparable interface.
func InsertionSortCustom(data Comparable) {
	length := data.Len()
	if length < 2 {
		return
	}

	for j := 1; j < length; j++ {
		for i := j - 1; i >= 0 && data.Compare(i+1, i) > 0; i-- {
			data.Swap(i+1, i)
		}
	}
}

// InsertionSort sorts a list of integers in ascending order.
func InsertionSort(elements ...int) []int {
	if elements == nil || len(elements) == 0 {
		return []int{}
	}

	for j := 1; j < len(elements); j++ {
		elem := elements[j]
		i := j - 1
		for ; i >= 0 && elements[i] > elem; i-- {
			elements[i+1] = elements[i]
		}
		elements[i+1] = elem
	}

	return elements
}

// SelectionSort sorts a list of integers in ascending order.
func SelectionSort(elements ...int) []int {
	if elements == nil || len(elements) == 0 {
		return []int{}
	}

	length := len(elements)

	for i := 0; i < length-1; i++ {
		smallest := i
		for j := i + 1; j < length; j++ {
			if elements[smallest] > elements[j] {
				smallest = j
			}
		}
		elements[i], elements[smallest] = elements[smallest], elements[i]
	}

	return elements
}

// MergeSort sorts a list of integers in ascending order.
func MergeSort(elements ...int) {
	if elements == nil || len(elements) == 0 {
		return
	}

	mergeSort(0, len(elements)-1, elements...)
}

func mergeSort(startIndex, endIndex int, elements ...int) {
	if startIndex < endIndex {
		mid := startIndex + (endIndex-startIndex)/2
		mergeSort(startIndex, mid, elements...)
		mergeSort(mid+1, endIndex, elements...)
		merge(startIndex, mid, endIndex, elements...)
	}
}

func merge(startIndex, midIndex, endIndex int, elements ...int) {
	firstPart := copy(startIndex, midIndex, elements)
	secondPart := copy(midIndex+1, endIndex, elements)

	for i, j, k := 0, 0, startIndex; k <= endIndex; k++ {
		if i >= len(firstPart) {
			elements[k] = secondPart[j]
			j++
		} else if j >= len(secondPart) {
			elements[k] = firstPart[i]
			i++
		} else if firstPart[i] <= secondPart[j] {
			elements[k] = firstPart[i]
			i++
		} else {
			elements[k] = secondPart[j]
			j++
		}
	}
}

func copy(start, end int, arr []int) []int {
	size := end - start + 1
	copy := make([]int, size)
	for i, j := start, 0; i <= end; i, j = i+1, j+1 {
		copy[j] = arr[i]
	}

	return copy
}
