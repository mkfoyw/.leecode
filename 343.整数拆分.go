/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

// @lc code=start
func integerBreak(n int) int {
	if n <= 0 {
		return 0
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = -1
	}

	var dfs func(int) int
	dfs = func(cur int) int {
		if dp[cur] != -1 {
			return dp[cur]
		}

		dp[cur] = cur - 1
		for i := 1; i < cur; i++ {
			dp[cur] = max(dp[cur], max(i*(cur-i), i*dfs(cur-i)))

		}
		return dp[cur]
	}
	return dfs(n)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

