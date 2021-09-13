/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lc code=start
type NumArray struct {
	nums []int
	sum  []int
}

var sum []int

func Constructor(nums []int) NumArray {
	na := NumArray{nums, make([]int, len(nums))}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			na.sum[i] = nums[i]
		} else {
			na.sum[i] = na.sum[i-1] + nums[i]
		}
	}
	return na
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.sum[right]
	} else {
		return this.sum[right] - this.sum[left-1]
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end

