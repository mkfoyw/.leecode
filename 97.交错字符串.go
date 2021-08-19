/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] 交错字符串
 */

// @lc code=start

func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n1+n2 != n3 {
		return false
	}

	dp := make([][]bool, 0)
	for i := 0; i <= n1; i++ {
		dp = append(dp, make([]bool, n2+1))
	}

	dp[0][0] = true
	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || dp[i-1][j] && s1[i-1] == s3[p]
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || dp[i][j-1] && s2[j-1] == s3[p]
			}
		}
	}
	return dp[n1][n2]
}

// Accepted
// 106/106 cases passed (732 ms)
// Your runtime beats 5.69 % of golang submissions
// Your memory usage beats 5.09 % of golang submissions (150 MB)

// var dp [][][][]int8
// func isInterleave(s1 string, s2 string, s3 string) bool {
// 	n1, n2, n3 := len(s1), len(s2), len(s3)
// 	dp = nil
// 	for i := 0; i <= n1; i++ {
// 		var t2 [][][]int8
// 		for j := 0; j <= n2; j++ {
// 			var t1 [][]int8
// 			for k := 0; k <= n3; k++ {
// 				t0 := make([]int8, 3)
// 				t1 = append(t1, t0)
// 			}
// 			t2 = append(t2, t1)
// 		}
// 		dp = append(dp, t2)
// 	}
// 	return dfs(s1, s2, s3, 1) == 1 || dfs(s1, s2, s3, 2) == 1
// }

// func dfs(s1, s2, s3 string, mod int) int8 {
// 	n1, n2, n3 := len(s1), len(s2), len(s3)
// 	if dp[n1][n2][n3][mod] != 0 {
// 		return dp[n1][n2][n3][mod]
// 	}

// 	if n1 == 0 && n2 == 0 && n3 == 0 {
// 		dp[n1][n2][n3][mod] = 1
// 		return 1
// 	}

// 	if (mod == 1 && n1 == 0 && n3 != 0) || (mod == 2 && n2 == 0 && n3 != 0) {
// 		dp[n1][n2][n3][mod] = -1
// 		return -1
// 	}

// 	for i := 1; i <= n3; i++ {
// 		if mod == 1 {
// 			if i <= n1 && s1[:i] == s3[:i] {
// 				if dfs(s1[i:], s2, s3[i:], 2) == 1 {
// 					dp[n1][n2][n3][mod] = 1
// 					return 1
// 				}
// 			}
// 		} else {
// 			if i <= n2 && s2[:i] == s3[:i] {
// 				if dfs(s1, s2[i:], s3[i:], 1) == 1 {
// 					dp[n1][n2][n3][mod] = 1
// 					return 1
// 				}
// 			}
// 		}
// 	}

// 	dp[n1][n2][n3][mod] = -1
// 	return -1
// }

// @lc code=end

