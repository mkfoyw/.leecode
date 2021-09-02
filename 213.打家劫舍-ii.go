/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
// 先考虑简单情况：
//     只有一家的情况
//     只有两家的情况
//     有两家以上的情况
//			 如果我们偷第一家的话， 那么我们最后一家不能够偷
//           如果我们不偷第一家的话， 那么我们可以偷最后一家

// dp[i][0]: 表示我们不偷第一家， 前 i 户能够获得的收益
// dp[i][1]: 表示我们偷第一家， 前 i 户能够获得收益
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	dp := make([][2]int, n, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = nums[i]
		} else if i == 1 {
			dp[i][0] = nums[1]
			dp[i][1] = nums[0]
		} else if i == n-1 {
			dp[i][0] = max(dp[i-1][0], dp[i-2][0]+nums[i])
			dp[i][1] = dp[i-1][1]
		} else {
			dp[i][0] = max(dp[i-1][0], dp[i-2][0]+nums[i])
			dp[i][1] = max(dp[i-1][1], dp[i-2][1]+nums[i])
		}
	}

	return max(dp[n-1][0], dp[n-1][1])
}

// @lc code=end

