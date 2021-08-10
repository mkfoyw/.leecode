package main

func longestValidParentheses(s string) int {
	n := len(s)
	dp := make([]int, n, n)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := 0
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i-2 >= 0 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else {
				if i-1-dp[i-1] >= 0 && s[i-dp[i-1]-1] == '(' {
					dp[i] = dp[i-1] + 2
					if i-dp[i] >= 0 {
						dp[i] = dp[i] + dp[i-dp[i]]
					}
				}

			}
		}

		res = max(res, dp[i])
	}
	return res
}
