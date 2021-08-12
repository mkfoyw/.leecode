package main

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	t1 := []int{0, 1, 0, 1}
	res := largestRectangleArea(t1)
	t.Log(res)

}
