/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
//  解题思路：
//  dp[i]表示前i个字符有多少种映射方案
//  第 i 个字符单独映射一个字符  dp[i-2]
//  第 i 个字符和第 i -1 个字符共同映射一个字符 dp[i-1]
//   dp[i] = dp[i-2] + dp[i-1]
func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if i == 1 {
			if s[i-1] == '0' {
				return 0
			} else {
				dp[i] = 1
			}
		} else {
			if s[i-1] == '0' {
				if s[i-2] != '0' && s[i-2]-'0' <= 2 {
					dp[i] = dp[i-2]
				} else {
					return 0
				}
			} else {
				if s[i-2] != '0' && (s[i-2]-'0')*10+s[i-1]-'0' <= 26 {
					dp[i] = dp[i-1] + dp[i-2]
				} else {
					dp[i] = dp[i-1]
				}
			}
		}
	}
	return dp[n]
}

// @lc code=end

