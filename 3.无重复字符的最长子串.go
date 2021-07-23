/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	// 用一个集合来位置到到s[..:i]最长的非重复子串
	mm := map[byte]bool{}

	mx := 0
	last := 0
	for i := 0; i < len(s); i++ {
		if mm[s[i]] == true {
			for j := last; j <= i; j++ {
				last = j + 1
				if s[i] != s[j] {
					delete(mm, s[j])
				} else {
					break
				}

			}
		}

		mm[s[i]] = true
		if len(mm) > mx {
			mx = len(mm)
		}
	}
	return mx

}

// @lc code=end

