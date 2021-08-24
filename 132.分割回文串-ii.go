/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] 分割回文串 II
 */

// @lc code=start
// dp[i][j]： 表示从字符串s[i,j]全部分割成回文串的最小分割次数
// dp[i][j] = max(dp[i][j], dp[i][i+1]+dp[i+1][j])
func minCut(s string) int {
	n := len(string)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)

	for i := 1; i < n; i++ {
		for j = 0; i+j <= n; j++ {
			for k := 0; k < i; k++ {

			}
		}
	}
}

// @lc code=end

