/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start

// dp[i][j] 表示前i个数选取一些数是否能够组成j
func canPartition(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return true
	}

	sum := 0
	max := nums[0]
	for i := 0; i < n; i++ {
		sum += nums[i]
		if max < nums[i] {
			max = nums[i]
		}
	}

	target := sum / 2
	if sum&1 == 1 || max > target {
		return false
	}

	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= target; j++ {
			if j >= nums[i-1] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][target]
}

// @lc code=end

