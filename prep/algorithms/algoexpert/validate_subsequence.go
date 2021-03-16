package main

// Time: O(n)
// Space: O(1)
func IsValidSubsequence(array []int, sequence []int) bool {
	seqId := 0
	arrId := 0

	for arrId < len(array) && seqId < len(sequence) {
		if array[arrId] == sequence[seqId] {
			seqId++
		}
		arrId++
	}

	return len(sequence) == seqId
}
