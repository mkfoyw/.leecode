/*
 * @lc app=leetcode.cn id=376 lang=golang
 *
 * [376] 摆动序列
 */

// @lc code=start

// up[i]: 表示前 i 个元素中某一个元素为尾元素的最长上升摆动序列
// down[i]： 表示前 i 个元素中某一个元素为尾元素的最长下降摆动序列

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	res := 1
	up, down := make([]int, n), make([]int, n)
	up[0], down[0] = 1, 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			up[i] = max(up[i-1], down[i-1]+1)
			down[i] = down[i-1]
		} else if nums[i] < nums[i-1] {
			up[i] = up[i-1]
			down[i] = max(down[i-1], up[i-1]+1)
		} else {
			up[i] = up[i-1]
			down[i] = down[i-1]
		}
		res = max(res, max(up[i], down[i]))
	}
	return res
}

// dp[i][0]：表示一nums[i]结尾，并且与前一个数的叉值为正数的最大长度
// dp[i][1]：表示一nums[i]结尾，并且与前一个数的叉值为负数的最大长度
func wiggleMaxLength2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	res := 1
	dp := make([][2]int, n)
	dp[0][0] = 1
	dp[0][1] = 1
	for i := 1; i < n; i++ {
		dp[i][0], dp[i][1] = 1, 1
		for j := 0; j < i; j++ {
			if nums[i]-nums[j] > 0 {
				dp[i][0] = max(dp[i][0], dp[j][1]+1)
			} else if nums[i]-nums[j] < 0 {
				dp[i][1] = max(dp[i][1], dp[j][0]+1)
			}
			res = max(res, max(dp[i][0], dp[i][1]))
		}
	}
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

