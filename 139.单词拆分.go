/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start

// dp[i]: 表示前i个字符是否能够满足题目条件的分割
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	if n == 0 {
		return true
	}

	flag := map[string]bool{}
	for _, word := range wordDict {
		flag[word] = true
	}

	dp := make([]bool, n)

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				dp[i] = dp[i] || flag[s[:i+1]]
			} else {
				dp[i] = dp[i] || (dp[j-1] && flag[s[j:i+1]])
			}
			if dp[i] == true {
				break
			}
		}
	}
	return dp[n-1]
}

// @lc code=end

