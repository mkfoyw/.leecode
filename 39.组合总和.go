/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
// 在搜索所有的可行解， 可以考虑 dfs 回溯
func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	if n == 0 {
		return [][]int{}
	}

	res := [][]int{}

	var dfs func(int, int)
	combine := []int{}

	dfs = func(idx, target int) {
		//找到一个可行的解
		if target == 0 {
			res = append(res, append([]int{}, combine...))
			return
		}
		if target < 0 || idx >= n {
			return
		}

		//记录状态
		curlen := len(combine)
		// 不选择
		dfs(idx+1, target)
		// 选择
		for i := 1; i*candidates[idx] <= target; i++ {
			combine = append(combine, candidates[idx])
			dfs(idx+1, target-i*candidates[idx])
		}
		//恢复状态
		combine = combine[:curlen]
	}
	dfs(0, target)
	return res
}

// @lc code=end
