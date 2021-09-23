/*
 * @lc app=leetcode.cn id=410 lang=golang
 *
 * [410] 分割数组的最大值
 */

// @lc code=start

//二分查找 + 贪心
// sum[j]: 表示前j个整数的和。 那么sum[n] 就是该题答案的上界。
// 然后我们[0, sum[n]]内进行二分搜索， 判断是否有符合问题的答案。

func splitArray(nums []int, m int) int {
	n := len(nums)
	if n < m {
		return -1
	}

	left, right := 0, 0
	for i := 0; i < n; i++ {
		left = max(left, nums[i])
		right += nums[i]
	}

	check := func(mid int) bool {
		group := 1
		add := 0

		for i := 0; i < n; i++ {
			if nums[i] > mid {
				return false
			}
			if add+nums[i] > mid {
				group++
				add = nums[i]
			} else {
				add += nums[i]
			}
		}

		return group <= m
	}

	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// dp[i][m] 表示前 i 个数字分成 m 组的最大值
func splitArray2(nums []int, m int) int {
	n := len(nums)
	if n < m {
		return -1
	}

	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = sum[n]
		}
	}

	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= min(i, m); j++ {
			for k := 0; k < i; k++ {
				dp[i][j] = min(dp[i][j], max(dp[k][j-1], sum[i]-sum[k]))
			}
		}
	}
	return dp[n][m]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

