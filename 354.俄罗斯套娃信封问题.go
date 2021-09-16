
/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 */

// @lc code=start

func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n == 0 {
		return 0
	}

	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] > envelopes[j][1]
	})

	height := []int{}
	for i := 0; i < n; i++ {
		if len(height) == 0 || height[len(height)-1] < envelopes[i][1] {
			height = append(height, envelopes[i][1])
			continue
		}
		//找到height 中第一个币envelopes[i][1] 大的数
		left, right := 0, len(height)
		for left < right {
			mid := left + (right-left)/2
			if height[mid] < envelopes[i][1] {
				left = mid + 1
			} else {
				right = mid
			}
		}
		height[left] = envelopes[i][1]
	}
	return len(height)
}

// @lc code=end
