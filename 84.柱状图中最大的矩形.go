/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */
// 暴力算法： 枚举可能矩形的高度或者高度可以暴力得出结果。
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

		if len(stack) == 0 {
			left[i] = i
			stack = append(stack, i)
			continue
		}

		cur := len(stack) - 1
		for cur >= 0 {
			if heights[stack[cur]] <= heights[i] {
				break
			}
			cur--
		}

		if cur < 0 {
			left[i] = stack[0]
			stack = stack[:0]
			stack = append(stack, i)
		} else {
			if heights[stack[cur]] == heights[i] {
				left[i] = stack[cur]
				stack = stack[:cur+1]
			} else {
				left[i] = stack[cur] + 1
				stack = stack[:cur+1]
				stack = append(stack, i)
			}
		}
	}

	//找到最右边到 i 之间矩形高度都大于heights[i]的矩形编号
	right := make([]int, n, n)
	stack = stack[:0]
	for i := n - 1; i >= 0; i-- {
		if len(stack) == 0 {
			right[i] = i
			stack = append(stack, i)
			continue
		}

		cur := len(stack) - 1
		for cur >= 0 {
			if heights[stack[cur]] <= heights[i] {
				break
			}
			cur--
		}

		if cur < 0 {
			right[i] = stack[0]
			stack = stack[:0]
			stack = append(stack, i)

		} else {
			if heights[stack[cur]] == heights[i] {
				right[i] = stack[cur]
				stack = stack[:cur+1]
			} else {
				right[i] = stack[cur] - 1
				stack = stack[:cur+1]
				stack = append(stack, i)

			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		res = max(res, (right[i]-left[i]+1)*heights[i])
	}

	return res
}

// @lc code=end

