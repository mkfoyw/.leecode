/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */

// @lc code=start
// 记忆化搜索
func combinationSum3(k int, n int) [][]int {
	if n == 0 || k == 0 {
		return [][]int{}
	}

	res := [][]int{}

	var dfs func(int, int)
	oneRes := []int{}
	dfs = func(idx, target int) {
		if target == 0 && len(oneRes) == k {
			res = append(res, append([]int{}, oneRes...))
		}
		if idx >= 10 || len(oneRes) >= k {
			return
		}

		dfs(idx+1, target)
		oneRes = append(oneRes, idx)
		dfs(idx+1, target-idx)
		oneRes = oneRes[:len(oneRes)-1]
	}
	dfs(1, n)
	return res
}

// @lc code=end

