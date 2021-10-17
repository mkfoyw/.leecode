/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

// @lc code=start

// dp[i][j]: 表示前text1 的前i个字符串和text2前j个字符串的最长公共子序列的长度
// if text1[i] == text[j]
func longestCommonSubsequence(text1 string, text2 string) int {
	t1, t2 := len(text1), len(text2)
	if t1 == 0 || t2 == 0 {
		return 0
	}

	dp := make([][]int, t1+1)
	for i := 0; i <= t1; i++ {
		dp[i] = make([]int, t2+1)
	}

	for i := 1; i <= t1; i++ {
		for j := 1; j <= t2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[t1][t2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

