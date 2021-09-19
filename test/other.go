package main

import "fmt"

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
//
func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	if n == 0 {
		return [][]int{}
	}

	res := [][]int{}

	var dfs func(int, int)

	cur := []int{}
	dfs = func(pos, target int) {
		fmt.Printf("pos : %d, target: %d\n", pos, target)
		if target < 0 || pos >= n {
			return
		}
		if target == 0 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, cur)
			return
		}

		curlen := len(cur)
		// 不选择
		dfs(pos+1, target)
		// 选择
		for i := 1; i*candidates[pos] <= target; i++ {
			cur = append(cur, candidates[pos])
			dfs(i+1, target-i*candidates[pos])
		}
		cur = cur[:curlen]
	}
	dfs(0, target)
	return res
}

// @lc code=end
