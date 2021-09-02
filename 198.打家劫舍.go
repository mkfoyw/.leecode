/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start

// dp[i]: 表示偷窃前i 家能够获得最大收益
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n, n)
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i] = nums[i]
		} else if i == 1 {
			dp[i] = max(nums[i], nums[i-1])
		} else {
			dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		}
	}
	return dp[n-1]
}

// dp[i][0]: 不偷窃第 i 家住户能够获得最大的收益
// dp[i][1]: 偷窃第 i 家住户能够获得最大的收益
func rob1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := [][]int{}
	for i := 0; i < n; i++ {
		dp = append(dp, make([]int, 2, 2))
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = nums[i]
		} else {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1])
			dp[i][1] = dp[i-1][0] + nums[i]
		}
	}
	return max(dp[n-1][0], dp[n-1][1])
}

// @lc code=end

