package main

// Time: O(2^n); Space: O(n) without tail recursion
/* func GetNthFib(n int) int {
	if n == 1 {
		return 0
	}

	if n == 2 {
		return 1
	}

	return GetNthFib(n-1) + GetNthFib(n-2)
}
*/

import "fmt"

// Memoized Solution
// Time: O(n); Space: O(n)
func GetNthFib(n int) int {
	memoized := map[int]int{1: 0, 2: 1}
	return callFib(n, memoized)
}

func callFib(n int, memoized map[int]int) int {
	if val, ok := memoized[n]; ok {
		fmt.Printf("Not memoized: %d\n", n)
		return val
	} else {
		memoized[n] = callFib(n-1, memoized) + callFib(n-2, memoized)
		return memoized[n]
	}
}
