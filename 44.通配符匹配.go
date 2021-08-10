/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] 通配符匹配
 */

// @lc code=start

// dp[i][j] 表示字符串p的前i个字符和字符串s的前j个字符串是否匹配
// p[i] 是普通字符
// p[i] 是'?' 的时候
// p[i] 是’*‘
func isMatch(s string, p string) bool {
	n, m := len(p), len(s)

	dp := make([][]bool, 0, n+1)
	for i := 0; i <= n; i++ {
		dp = append(dp, make([]bool, m+1))
	}

	dp[0][0] = true
	//注意字符串为空的时候
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[i][0] = true
		} else {
			break
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if p[i-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[i-1] == '?' || p[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[n][m]

}

// @lc code=end

