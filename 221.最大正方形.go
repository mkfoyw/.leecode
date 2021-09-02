/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */

// @lc code=start
// 这不是那个柱状矩形的问题吗
func maximalSquare(matrix [][]byte) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	if m == 0 {
		return 0
	}

	//统计矩形中每个元素往左累积的1个数
	flag := [][]int{}
	for i := 0; i < n; i++ {
		flag = append(flag, make([]int, m, m))
		cur += 0 //当前累积的1的个数
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

		for j : = 0; j < n; j ++{
			right[j] = n 
		}

		for j := 0; j < n; j++ {
			if len(stack) > 0 && flag[stack[len(stack)-1]][i] > flag[j][i] {
				right[stack[len(stack)-1]] = j
				stack =   stack[:len(stack)-1]
			}

			if len(stack) == {
				left[j] = -1 
			}else{
				left[j] = stack[len(stack)-1] 
			}
			stack = append(stack, flag[j][i])
		}

		for j := 0; j < n; j++{
			res = max(res, (right[i]-left[i]-1)*flag[j][i])
		}


	}

}

// @lc code=end

