/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 */

// @lc code=start
// dp[i][k] 表示字符串前s的前i个字符能够凑出多少种字符串t的前k个字符
// 当s[i-1] == t[k-1] 时,  dp[i][k] = dp[i-1][k-1] + dp[i-1][k]
// 当s[i-1] != k[k-1] 时,  dp[i][k] = dp[i-1][k]
func numDistinct(s string, t string) int {
	ns, nt := len(s), len(t)

	if ns < nt {
		return 0
	}

	dp := make([][]int, 0)
	for i := 0; i <= ns; i++ {
		dp = append(dp, make([]int, nt+1))
		dp[i][0] = 1
	}
	for j := 1; j <= nt; j++ {
		for i := 1; i <= ns; i++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[ns][nt]
}

// @lc code=end

