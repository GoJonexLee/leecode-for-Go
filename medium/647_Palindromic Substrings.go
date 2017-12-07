package medium

func countSubstrings(s string) int {

	n, res := len(s), 0
	dp := make([][]bool, len(s))
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if (s[i] == s[j]) && (j - i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
			}
			if dp[i][j] {
				res++
			}
		}
	}

	return res
}