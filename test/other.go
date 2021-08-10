package main

func isMatch(s string, p string) bool {
	n, m := len(p), len(s)

	dp := make([][]bool, 0, n+1)
	for i := 0; i <= n; i++ {
		dp = append(dp, make([]bool, m+1))
	}

	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[i][0] = true
		} else {
			break
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if p[i-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[i-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[n][m]

}