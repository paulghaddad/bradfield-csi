package main

func BubbleSort(array []int) []int {
	swapped := true

	// Since after each iteration the last-most element will be in its
	// final position, we can reduce how far along the array we need to traverse each time
	ceiling := 0

	for swapped {
		swapped = false

		for i, j := 0, 1; j < len(array)-ceiling; i, j = i+1, j+1 {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
				swapped = true
			}

		}
		ceiling++
	}

	return array
}
