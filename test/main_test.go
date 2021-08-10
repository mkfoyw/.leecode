package main

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	t1 := "(()())"
	res := longestValidParentheses(t1)
	t.Log(res)

}
