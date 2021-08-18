/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */
// 暴力算法： 枚举可能矩形的高度或者高度可以暴力得出结果。
//
//
// 在枚举高度的方法
// 首先我们悬着以选择某一个的柱子 i 的高度作为 h = heights[i]
// 然后向左寻在第一个高度小于 h 的柱子 j
// 然后向右寻找第一个高度小于 h 的柱子 k
// 这样在 j 和 k 之间的柱子都大于等于  h

// @lc code=start

func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	left := make([]int, n)
	right := make([]int, n)
	stack := []int{}
	for i := 0; i < n; i++ {
		right[i] = n
	}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	res := 0
	for i := 0; i < n; i++ {
		res = max(res, (right[i]-left[i]-1)*heights[i])
	}
	return res
}

// @lc code=end

