/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
// dp[i][target]: 表示前i个数之后能够组成target 的数目

func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		dp[i] = map[int]int{}
	}

	var dfs func(int, int) int

	dfs = func(cur int, target int) int {
		if cur == n && target == 0 {
			return 1
		}
		if cur >= n {
			return 0
		}

		if _, ok := dp[cur][target]; ok {
			return dp[cur][target]
		}

		res := 0
		res += dfs(cur+1, target-nums[cur])
		res += dfs(cur+1, target+nums[cur])

		dp[cur][target] = res
		return dp[cur][target]
	}
	return dfs(0, target)

}

// @lc code=end

