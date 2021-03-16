package main

// Time: O(n^2); Space: O(1)
func InsertionSort(array []int) []int {

	// Partition array into sorted and unsorted regions
	//for partition := 0; partition < len(array)-1; partition++ {
	for i := 1; i < len(array); i++ {

		// Take first element from unsorted region and swap it with
		// as many of the sorted ones as needed
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}

	return array
}
