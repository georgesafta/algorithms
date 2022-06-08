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
