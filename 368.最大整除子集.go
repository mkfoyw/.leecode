/*
 * @lc app=leetcode.cn id=368 lang=golang
 *
 * [368] 最大整除子集
 */

// @lc code=start
// 整除具有传递性 a|b, b|c -> a |c

func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	mm := make([]int, n)
	pre := make([]int, n)

	mx := -1
	mp := -1
	for i := 0; i < n; i++ {
		mm[i] = 1
		pre[i] = -1
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				if mm[i] < mm[j]+1 {
					mm[i] = mm[j] + 1
					pre[i] = j
				}
			}
		}
		if mx < mm[i] {
			mx = mm[i]
			mp = i
		}
	}

	res := []int{}
	for mp != -1 {
		res = append(res, nums[mp])
		mp = pre[mp]
	}
	return res
}

// @lc code=end

