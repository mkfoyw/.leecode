/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])

	dp := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n, n))
	}

	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				if obstacleGrid[i][j] == 1 {
					dp[i][j] = 0
				} else {
					if i > 0 {
						dp[i][j] = dp[i-1][j]
					} else if j > 0 {
						dp[i][j] = dp[i][j-1]
					}
				}
			} else {
				if obstacleGrid[i][j] == 1 {
					dp[i][j] = 0
				} else {
					dp[i][j] = dp[i-1][j] + dp[i][j-1]
				}
			}
		}
	}
	return dp[m-1][n-1]
}

// @lc code=end

