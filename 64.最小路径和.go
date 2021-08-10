/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	dp := make([][]int, 0, m)
	for j := 0; j < m; j++ {
		dp = append(dp, make([]int, n, n))
	}

	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				if i > 0 {
					dp[i][j] = dp[i-1][j] + grid[i][j]
				} else if j > 0 {
					dp[i][j] = dp[i][j-1] + grid[i][j]
				}
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}

// @lc code=end

