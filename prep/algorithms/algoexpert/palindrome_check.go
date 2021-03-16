package main

import "bytes"

// Pointer technique
// O(n) time; O(1) space
func IsPalindrome1(str string) bool {
	left := 0
	right := len(str) - 1

	for left < right {
		if str[left] != str[right] {
			return false
		}

		left++
		right--
	}

	return true
}

// Recursive
// O(n) time; O(n) space
func IsPalindrome2(str string) bool {
	strLen := len(str)
	if strLen <= 1 {
		return true
	}

	return str[0] == str[strLen-1] && IsPalindrome(str[1:strLen-1])
}

// Compare string to reversed string
// O(n) time; O(n) space
func IsPalindrome3(str string) bool {
	var buf bytes.Buffer
	for i := len(str) - 1; i >= 0; i-- {
		buf.WriteByte(str[i])
	}

	return str == buf.String()
}
