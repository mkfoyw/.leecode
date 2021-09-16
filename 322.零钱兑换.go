/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// dp[i]： 表示兑换 i 需要的最少数， 那么可以得到
// dp[i] = min(dp[i-coin[0]], dp[i-coin[1]], ...., dp[i-coin[k]])

// @lc code=start
func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}

	// -2 表示未经处理， -1 表示结果不存在
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = -2
	}
	dp[0] = 0

	var dfs func(int) int

	dfs = func(cur int) int {
		if cur < 0 {
			return -1
		}
		if dp[cur] != -2 {
			return dp[cur]
		}

		res := -1
		for i := 0; i < len(coins); i++ {
			tres := dfs(cur - coins[i])
			if tres == -1 {
				continue
			}

			if res == -1 || res > tres+1 {
				res = tres + 1
			}
		}
		dp[cur] = res
		return dp[cur]
	}
	return dfs(amount)
}

// @lc code=end

