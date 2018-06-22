package medium

func longestPalindrome(s string) string {
	cp := len(s)

	dp := make([][]bool, cp)
	for i := 0; i < cp; i++ {
		dp[i] = make([]bool, cp)
	}

	left, len := 0, 1
	for i := 0; i < cp; i++ {
		for j := 0; j < i; j++ {
			if (s[i] == s[j]) && (i-j < 2 || dp[j+1][i-1]) {
				dp[j][i] = true
			}

			if dp[j][i] && len < i-j+1 {
				len = i - j + 1
				left = j
			}
		}
		dp[i][i] = true
	}
	return string(s[left : left+len])
}

func longestPalindrome(s string) string {
	length := len(s)
	find := func(left, right int) string {
		for left >= 0 && right < length && s[left] == s[right] {
			left--
			right++
		}
		return s[left+1 : right]
	}
	ret := ""
	for i := 0; i < length; i++ {
		if p := find(i-1, i+1); len(p) > len(ret) {
			ret = p
		}
		if i < length-1 {
			if p := find(i, i+1); len(p) > len(ret) {
				ret = p
			}
		}
	}
	return ret
}
