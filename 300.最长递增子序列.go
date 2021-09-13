/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start

//dp[i]:表示长度为 i 递增数列的最后一个值
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	dp := []int{}
	for i := 0; i < n; i++ {
		if len(dp) == 0 {
			dp = append(dp, nums[i])
		} else if nums[i] > dp[len(dp)-1] {
			dp = append(dp, nums[i])
		} else {
			left, right := 0, len(dp)
			//找到第一个大于 nums[i] 的值
			for left < right {
				mid := left + (right-left)/2
				if dp[mid] < nums[i] {
					left = mid + 1
				} else {
					right = mid
				}
			}
			dp[left] = nums[i]
		}
	}
	return len(dp)

}

// dp[i]: 表示以数字 nums[i]结尾的最长递增子序列的长度
func lengthOfLIS2(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	dp := make([]int, n, n)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := 0
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(dp[i], res)
	}
	return res

}

// @lc code=end

