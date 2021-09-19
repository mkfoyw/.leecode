package main

/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start

// dp[i][j]： 表示字符串s前i个字符串是否是字符串t前j个字符串的子序列
func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	if n == 0 || m == 0 {
		if n == 0 {
			return true
		}
		if m == 0 {
			return false
		}
	}

	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}
	for i := 0; i <= m; i++ {
		dp[0][i] = true
	}

	for i := 1; i <= n; i++ {
		for j := i; j <= m; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] || dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[n][m]
}

// @lc code=end
