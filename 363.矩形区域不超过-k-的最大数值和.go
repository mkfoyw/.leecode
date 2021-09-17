/*
 * @lc app=leetcode.cn id=363 lang=golang
 *
 * [363] 矩形区域不超过 K 的最大数值和
 */

// @lc code=start

// 关键是枚举矩形
// 然后枚举矩形的上边界 和 下边界
// 然后在枚举矩形的左边界和右边界
// 然后找到一个  Sr - Sl <= k 的最大值 =》 可以转化位 Sl >= Sr -k
// 因此我们需要维持一个递增的 sl 序列

func maxSumSubmatrix(matrix [][]int, k int) int {
	row, col := len(matrix), len(matrix[0])
	sum := make([][]int, row)
	for i := 0; i < row; i++ {
		sum[i] = make([]int, col)
		cur := 0
		for j := 0; j < col; j++ {
			cur += matrix[i][j]
			if i == 0 {
				sum[i][j] = cur
			} else {
				sum[i][j] = sum[i-1][j] + cur
			}
		}
	}

	res := -10000
	for r1 := 0; r1 < row; r1++ {
		for c1 := 0; c1 < col; c1++ {
			for r2 := r1; r2 < row; r2++ {
				for c2 := c1; c2 < col; c2++ {
					cur := sum[r2][c2]
					if c1 != 0 {
						cur -= sum[r2][c1-1]
					}
					if r1 != 0 {
						cur -= sum[r1-1][c2]
					}
					if r1 != 0 && c1 != 0 {
						cur += sum[r1-1][c1-1]
					}
					if cur <= k {
						if cur > res {
							res = cur
						}
					}
				}
			}
		}
	}
	return res
}

// @lc code=end

