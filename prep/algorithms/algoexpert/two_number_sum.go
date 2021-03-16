package main

// time: O(n)
// space: O(n)
func TwoNumberSum(array []int, target int) []int {
	numSet := make(map[int]bool)

	for _, num := range array {
		complement := target - num
		if _, present := numSet[complement]; present {
			return []int{num, complement}
		}

		numSet[num] = true
	}

	return []int{}
}
