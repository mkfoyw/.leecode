package main

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	begin := 0
	maxLen := 0

	for i := 0; i < len(s); i++ {
		oddLen := expandAroundCenter(s, i, i)
		evenLen := expandAroundCenter(s, i, i+1)

		curMax := 0
		if oddLen < evenLen {
			curMax = evenLen
		} else {
			curMax = oddLen
		}

		if maxLen < curMax {
			maxLen = curMax
			begin = i - (maxLen-1)/2
		}
	}

	return string(s[begin : begin+maxLen])
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		} else {
			left--
			right++
		}
	}
	return right - left - 1
}
