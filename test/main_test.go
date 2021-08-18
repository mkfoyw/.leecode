package main

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	s1 := "ccabcbabcbabbbbcbb"
	s2 := "bbbbabccccbbbabcba"
	res := isScramble(s1, s2)
	t.Log(res)
}
