/*
 * @lc app=leetcode.cn id=377 lang=golang
 *
 * [377] 组合总和 Ⅳ
 */

// @lc code=start
// 记忆化搜索
func combinationSum4(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, target+1)
	for i := 0; i <= target; i++ {
		dp[i] = -1
	}
	dp[0] = 1

	var dfs func(int) int
	dfs = func(target int) int {
		if target >= 0 && dp[target] != -1 {
			return dp[target]
		}
		if target < 0 {
			return 0
		}

		cur := 0
		for i := 0; i < n; i++ {
			cur += dfs(target - nums[i])
		}
		dp[target] = cur
		return dp[target]
	}
	dfs(target)
	return dp[target]

}

// @lc code=end

