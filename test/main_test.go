package main

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {

	res := calculateMinimumHP([][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	})
	t.Log(res)
}
