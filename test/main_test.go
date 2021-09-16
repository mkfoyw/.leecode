package main

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {

	res := maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}})
	t.Log(res)
}
