package main

/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */

// @lc code=start

func maximalSquare(matrix [][]byte) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}

	m := len(matrix[0])
	if m == 0 {
		return 0
	}

	dp := [][]int{}
	for i := 0; i < n; i++ {
		dp = append(dp, make([]int, m))
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	//dp[i][j] 表示以matrix[i][j]为左下角最大的矩形变长
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
			} else if j == 0 || i == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
			res = max(res, dp[i][j])
		}
	}
	return res * res
}

// 这不是那个柱状矩形的问题吗
func maximalSquare2(matrix [][]byte) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	if m == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	//统计矩形中每个元素往左累积的1个数
	flag := [][]int{}
	for i := 0; i < n; i++ {
		flag = append(flag, make([]int, m, m))
		cur := 0 //当前累积的1的个数
		for j := 0; j < m; j++ {
			if matrix[i][j] == '0' {
				cur = 0
				flag[i][j] = 0
			} else {
				cur += 1
				flag[i][j] = cur
			}

		}
	}

	res := 0
	for i := 0; i < m; i++ {
		left, right := make([]int, n), make([]int, n)
		stack := []int{}

		for j := 0; j < n; j++ {
			right[j] = n
		}

		for j := 0; j < n; j++ {
			for len(stack) > 0 && flag[stack[len(stack)-1]][i] >= flag[j][i] {
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

		for j := 0; j < n; j++ {
			t := min(right[j]-left[j]-1, flag[j][i])
			res = max(res, t*t)
		}

	}
	return res

}

// @lc code=end
