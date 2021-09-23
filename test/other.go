package main

/*
 * @lc app=leetcode.cn id=446 lang=golang
 *
 * [446] 等差数列划分 II - 子序列
 */

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	res := 0
	mp := map[int]map[int]int{}
	for i := 0; i < n; i++ {
		mp[i] = map[int]int{}
		for j := 0; j < i; j++ {
			d := nums[i] - nums[j]
			res += mp[j][d]
			mp[i][d] += mp[j][d] + 1
		}
	}
	return res
}

// @lc code=end
