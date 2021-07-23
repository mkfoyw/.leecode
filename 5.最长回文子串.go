/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start

// 中心扩展法：
// 如果回文长度为奇数， 那么回文的中心只有一个 left == right
// 如果回文长度为偶数， 那么回文的中心有两个， right = left+1

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

// 动态规划法
func longestPalindrome1(s string) string {
	if len(s) == 0 {
		return ""
	}
	// dp[i][j] 表示以s[i]到s[j]是否是一个回文字符串
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	//初始化dp数组
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}

	for i := 2; i <= len(s); i++ {
		for j := 0; j+i-1 < len(s); j++ {
			if s[j] == s[j+i-1] && (i == 2 || dp[j+1][j+i-2]) {
				dp[j][j+i-1] = true

			}
		}
	}

	for i := len(s); i > 0; i-- {
		for j := 0; j+i-1 < len(s); j++ {
			if dp[j][j+i-1] {
				return string(s[j : j+i])
			}
		}
	}
	return ""
}

// @lc code=end

