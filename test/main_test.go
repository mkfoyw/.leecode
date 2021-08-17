package main

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	t1 := []int{2, 1, 5, 6, 2, 3}
	res := largestRectangleArea(t1)
	t.Log(res)

}
