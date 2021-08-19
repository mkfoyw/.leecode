/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start

// dp[i][i]: 表示走到该位置的最小路径和
// dp[i][j] = min(dp[i-1][j],dp[i-1][j-1]) + triangle[i][j]
func minimumTotal(triangle [][]int) int {
	row := len(triangle)
	if row == 0 {
		return 0
	}
	dp := make([][]int, 0)
	for i := 0; i < row; i++ {
		dp = append(dp, make([]int, len(triangle[i])+1))
	}

	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}

	dp[0][0] = triangle[0][0]
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}

	res := dp[row-1][0]
	for i := 0; i < len(triangle[row-1]); i++ {
		res = min(res, dp[row-1][i])
	}
	return res
}

// @lc code=end

