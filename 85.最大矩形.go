/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */

// @lc code=start
// 在进行计算前， 我们先统计矩形中每个元素向左连续1的数量， 并记录在dp数组中。
// dp[i][j] 表示第 i 行第 j 列向左连续的1的个数。
// 因此我们可以枚举以 martix[i][j]为全是1的矩形左下行。
func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}

	dp := make([][]int, m, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				if matrix[i][j] == '0' {
					dp[i][j] = 0
				} else {
					dp[i][j] = 1
				}
			} else {
				if matrix[i][j] == '0' {
					dp[i][j] = 0
				} else {
					dp[i][j] = dp[i][j-1] + 1
				}
			}
		}
	}

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		left := make([]int, m)
		right := make([]int, m)
		stack := make([]int, 0)
		for j := 0; j < m; j++ {
			right[j] = m
		}
		for j := 0; j < m; j++ {
			for len(stack) > 0 && dp[stack[len(stack)-1]][i] >= dp[j][i] {
				right[stack[len(stack)-1]] = j
				stack = stack[:len(stack)-1]
			}

			if len(stack) == 0 {
				left[j] = -1
			} else {
				left[j] = stack[len(stack)-1]
			}
			stack = append(stack, j)
		}

		for j := 0; j < m; j++ {
			res = max(res, (right[j]-left[j]-1)*dp[j][i])
		}
	}
	return res
}

// @lc code=end


