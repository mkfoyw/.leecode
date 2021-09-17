/*
 * @lc app=leetcode.cn id=375 lang=golang
 *
 * [375] 猜数字大小 II
 */

// @lc code=star

// dp[j][j]: 表示在[i,j] 范围内猜中答案的最小的花费
var dp [][]int

func getMoneyAmount(n int) int {
	dp = make([][]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	return dfs(1, n)
}

// 表示答案在 [low, high]区间我们能够猜中答案的最小花费
func dfs(low, high int) int {
	if low >= high {
		return 0
	}
	if dp[low][high] != 0 {
		return dp[low][high]
	}

	mmres := 100000
	for i := low + (high-low)/2; i <= high; i++ {
		mmres = min(mmres, i+max(dfs(low, i-1), dfs(i+1, high)))
	}
	dp[low][high] = mmres

	return dp[low][high]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

