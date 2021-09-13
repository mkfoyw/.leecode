/*
 * @lc app=leetcode.cn id=304 lang=golang
 *
 * [304] 二维区域和检索 - 矩阵不可变
 */

// @lc code=start
// sum[i][j]： 表示以(0,0)为左上角, (i,j) 为右下角构成的矩形的和
type NumMatrix struct {
	matrix [][]int
	sum    [][]int
}

func Constructor(matrix [][]int) NumMatrix {
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
	return NumMatrix{matrix, sum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	res := this.sum[row2][col2]
	if col1 != 0 {
		res -= this.sum[row2][col1-1]
	}
	if row1 != 0 {
		res -= this.sum[row1-1][col2]
	}
	if col1 != 0 && row1 != 0 {
		res += this.sum[row1-1][col1-1]
	}
	return res
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end

