/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
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

