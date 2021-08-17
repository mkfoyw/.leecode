/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */
// 暴力算法： 枚举可能矩形的高度或者高度可以暴力得出结果。
//
//
//

// @lc code=start

func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	//从i开始， 从右往左找到第一个小于 height[i]的矩形编号
	left := make([]int, n, n)
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	right := make([]int, n, n)
	stack = stack[:0]
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	res := 0
	for i := 0; i < n; i++ {
		res = max(res, (right[i]-left[i]-1)*heights[i])
	}

	return res
}

// @lc code=end

