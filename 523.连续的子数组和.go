/*
 * @lc app=leetcode.cn id=523 lang=golang
 *
 * [523] 连续的子数组和
 */

// @lc code=start
// 同余定理 (a - b) % k == 0, 那么有 a %k == b%k
func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}

	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	mp := map[int]int{}
	for i := 0; i <= n; i++ {
		remainder := sum[i] % k
		if index, ok := mp[remainder]; ok {
			if i-index >= 2 {
				return true
			}
		} else {
			mp[remainder] = i
		}
	}
	return false
}

// @lc code=end

