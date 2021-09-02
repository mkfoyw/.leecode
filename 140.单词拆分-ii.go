/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 */

// @lc code=start
// dp[i]：表示前i个字符满足题目条件有多少种切割方案
func wordBreak(s string, wordDict []string) []string {
	n := len(s)
	if n == 0 {
		return []string{}
	}

	flag := map[string]bool{}
	for _, word := range wordDict {
		flag[word] = true
	}

	dp := [][]string{}
	for i := 0; i < n; i++ {
		dp = append(dp, make([]string, 0))
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if flag[s[j:i+1]] {
				if j == 0 {
					dp[i] = append(dp[i], s[j:i+1])
				} else {
					for _, cur := range dp[j-1] {
						dp[i] = append(dp[i], cur+" "+s[j:i+1])
					}
				}
			}
		}
	}
	return dp[n-1]
}

// @lc code=end

