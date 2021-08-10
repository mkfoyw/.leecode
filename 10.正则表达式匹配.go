/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
// dp[i][j] 表示模式p的前i个字符是否于字符串s的前j个字符串是否匹配
// 当 p[i] 是普通字符或者 .时：
//	     如果  dp[i-1][j-1] = true, 且 p[i] == s[i]，
//       那么dp[i][j] = true, 否则 dp[i][j] = false

// 当 p[i] 是 * 字符时
//		如果 dp[i-1][0-j]存在一个为true时，
//      那么dp[i][j] = true, 否则

func isMatch(s string, p string) bool {
	m, n := len(p), len(s)

	matches := func(i, j int) bool {
		if j < 0 {
			return false
		}
		if p[i] == '.' {
			return true
		}
		if p[i] == s[j] {
			return true
		}
		return false
	}

	dp := make([][]bool, 0, m+1)
	for i := 0; i <= m; i++ {
		dp = append(dp, make([]bool, n+1, n+1))
	}
	dp[0][0] = true
	for i := 1; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if p[i-1] == '*' {
				// 匹配0次
				dp[i][j] = dp[i-2][j]

				// 匹配一次
				if matches(i-2, j-1) {
					dp[i][j] = dp[i][j] || dp[i][j-1]
				}
			} else if matches(i-1, j-1) {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]

}

// @lc code=end

