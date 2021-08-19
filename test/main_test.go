package main

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	triange := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	res := minimumTotal(triange)
	t.Log(res)
}
