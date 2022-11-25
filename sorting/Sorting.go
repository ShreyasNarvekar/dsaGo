package sorting

func BubbleSort(a []int) {
	// {12, 6, 54, 9, 18}
	for j := len(a) - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
}

func SelectionSort(array []int) {
	for i := 0; i < len(array); i++ {
		minIndex := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			array[i], array[minIndex] = array[minIndex], array[i]
		}
	}
}
