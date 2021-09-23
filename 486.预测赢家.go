/*
 * @lc app=leetcode.cn id=486 lang=golang
 *
 * [486] 预测赢家
 */

// @lc code=start
// dp[i][j]： 表示在nums[i][j] 开始选择， 能够获得最大的得分。

func PredictTheWinner(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return true
	}

	sum := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			sum[i] = nums[i]
		} else {
			sum[i] = sum[i-1] + nums[i]
		}
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	var dfs func(int, int) int

	dfs = func(i, j int) int {
		if i > j {
			return 0
		}

		if dp[i][j] != -1 {
			return dp[i][j]
		}

		cur := 0
		if i == 0 {
			cur = sum[j]
		} else {
			cur = sum[j] - sum[i-1]
		}

		cur = cur - min(dfs(i+1, j), dfs(i, j-1))
		dp[i][j] = cur
		return dp[i][j]
	}
	res := dfs(0, n-1)
	return sum[n-1]-res <= res

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

