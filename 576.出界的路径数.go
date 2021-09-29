/*
 * @lc app=leetcode.cn id=576 lang=golang
 *
 * [576] 出界的路径数
 */

// @lc code=start

// dp[i][j][k]: 表示走k步到第i行和第j列有多少种方法

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

const MOD int = 1000000007

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	if m == 0 || n == 0 {
		return 0
	}

	m1 := make([][]int, m)
	for i := 0; i < m; i++ {
		m1[i] = make([]int, n)
	}
	m2 := make([][]int, m)
	for i := 0; i < m; i++ {
		m2[i] = make([]int, n)
	}

	m1[startRow][startColumn] = 1

	res := 0
	for step := 1; step <= maxMove; step++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if m1[i][j] != 0 {
					for _, dir := range directions {
						if i+dir[0] >= 0 && i+dir[0] < m && j+dir[1] >= 0 && j+dir[1] < n {
							m2[i+dir[0]][j+dir[1]] = (m2[i+dir[0]][j+dir[1]] + m1[i][j]) % MOD
						} else {
							res = (res + m1[i][j]) % MOD
						}
					}
				}
			}
		}

		t1 := m1
		m1 = m2
		m2 = t1
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				m2[i][j] = 0
			}
		}

	}
	return res
}

// @lc code=end

