package main

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	testCace := "cbbd"
	res := longestPalindrome(testCace)
	t.Log(res)
	if len(res) == 0 {
		return
	}
}
