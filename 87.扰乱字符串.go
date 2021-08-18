/*
 * @lc app=leetcode.cn id=87 lang=golang
 *
 * [87] 扰乱字符串
 */

// @lc code=start
// 暴力算法： 枚举字符扰乱后的所有形式， 如果存在一种形式和目标字符传相等， 那么返回true， 否则返回false
// dp[i][j][k] 表示字符串s1[i:i+k] 和 s2[j:j+k] 是否是和谐的字符串

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var dp [][][]int8
	n := len(s1)
	for i := 0; i < n; i++ {
		t1 := make([][]int8, 0)
		for j := 0; j < n; j++ {
			t2 := make([]int8, n+1)
			for k := 0; k < n+1; k++ {
				t2[k] = -1
			}
			t1 = append(t1, t2)

		}
		dp = append(dp, t1)
	}

	var dfs func(i1, i2, length int) int8
	dfs = func(p1, p2, length int) int8 {
		//结果已经被计算
		if dp[p1][p2][length] != -1 {
			return dp[p1][p2][length]
		}

		if s1[p1:p1+length] == s2[p2:p2+length] {
			dp[p1][p2][length] = 1
			return 1
		}

		//字符串频繁性统计
		freq := [26]int{}
		for i := 0; i < length; i++ {
			freq[s1[p1+i]-'a']++
			freq[s2[p2+i]-'a']--
		}
		for _, val := range freq {
			if val != 0 {
				dp[p1][p2][length] = -1
				return -1
			}
		}

		for i := 1; i < length; i++ {
			// 不交换
			if dfs(p1, p2, i) == 1 && dfs(p1+i, p2+i, length-i) == 1 {
				dp[p1][p2][length] = 1
				return 1
			}
			//交换
			if dfs(p1, p2+length-i, i) == 1 && dfs(p1+i, p2, length-i) == 1 {
				dp[p1][p2][length] = 1
				return 1
			}
		}
		dp[p1][p2][length] = 0
		return 0
	}
	return dfs(0, 0, n) == 1
}

// func dfs(s1 string, s2 string) bool {
// 	if len(s1) != len(s2) {
// 		return false
// 	}
// 	if s1 == s2 {
// 		return true
// 	}

// 	n := len(s1)
// 	for i := 1; i < n; i++ {
// 		// 不交换
// 		if dfs(string(s1[:i]), string(s2[:i])) && dfs(string(s1[i:]), string(s2[i:])) {
// 			return true
// 		}
// 		// 交换
// 		if dfs(string(s1[:i]), string(s2[n-i:])) && dfs(string(s1[i:]), string(s2[:n-i])) {
// 			return true
// 		}
// 	}
// 	return false
// }

// @lc code=end
