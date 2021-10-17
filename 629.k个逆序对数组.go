/*
 * @lc app=leetcode.cn id=629 lang=golang
 *
 * [629] K个逆序对数组
 */

// @lc code=start
func kInversePairs(n int, k int) int {
	if k == 0 {
		return 0
	}

	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			for h := 1; h <= j; h++ {
				dp
			}
		}
	}
}

// @lc code=end

