package main

// Time: O(n^2)) | Space: O(1)
func SelectionSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		minIdx := i

		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minIdx] {
				minIdx = j
			}
		}

		swap(&array, i, minIdx)
	}

	return array
}

func swap(arrPtr *[]int, i, j int) {
	arr := *arrPtr

	arr[i], arr[j] = arr[j], arr[i]
}
