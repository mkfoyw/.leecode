package main

/*
 * @lc app=leetcode.cn id=174 lang=golang
 *
 * [174] 地下城游戏
 */

// @lc code=start

// 这有个问题，
func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 {
		return 0
	}
	n, m := len(dungeon), len(dungeon[0])

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

	var dfs func(int, int, int, int)
	res := -111111111
	dfs = func(i, j, need, left int) {
		if i == n-1 && j == m-1 {
			res = max(need, res)
			return
		}
		//往右走
		if j+1 < m {
			dfs(i, j+1, min(need, left+dungeon[i][j+1]), left+dungeon[i][j+1])
		}
		//往下走
		if i+1 < n {
			dfs(i+1, j, min(need, left+dungeon[i+1][j]), left+dungeon[i+1][j])
		}
	}

	dfs(0, 0, min(0, dungeon[0][0]), dungeon[0][0])
	return -res
}

// @lc code=end
