/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
// dp[i]: 表示组成数字i最少需要的平方数的个数

func numSquares(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1, n+1)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j])
		}
		dp[i] += 1
	}
	return dp[n]
}

func numSquares2(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = -1
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var dfs func(n int) int
	dfs = func(n int) int {
		if dp[n] != -1 {
			return dp[n]
		}
		res := n
		for i := 1; i*i <= n; i++ {
			res = min(res, dfs(n-i*i)+1)
		}
		dp[n] = res
		return dp[n]
	}

	dfs(n)
	return dp[n]
}

// @lc code=end

