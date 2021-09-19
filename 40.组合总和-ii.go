/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	freq := [][]int{}

	for i := 0; i < len(candidates); i++ {
		if i == 0 || candidates[i] != candidates[i-1] {
			freq = append(freq, []int{candidates[i], 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	res := [][]int{}
	var dfs func(int, int)
	oneRes := []int{}
	dfs = func(idx, target int) {
		//找到一个可行解
		if target == 0 {
			res = append(res, append([]int{}, oneRes...))
			return
		}
		if target < 0 || idx >= len(freq) {
			return
		}

		curLen := len(oneRes)
		//不选择
		dfs(idx+1, target)

		for i := 1; i <= freq[idx][1]; i++ {
			oneRes = append(oneRes, freq[idx][0])
			dfs(idx+1, target-i*freq[idx][0])
		}
		//恢复状态
		oneRes = oneRes[:curLen]
	}
	dfs(0, target)
	return res
}

// @lc code=end

