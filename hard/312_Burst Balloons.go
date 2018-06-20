package hard

func maxCoins(nums []int) int {
	var f = func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	nms := add(nums)
	n, dp := len(nums), initArrays(len(nums)) // dp[i][j] 表示打爆区间[i,j]中的气球可以得到的最多的金币
	for i := 1; i <= n; i++ {                 // 共需要打n次气球
		for left := 1; left <= n-i+1; left++ {
			right := left + i - 1
			for k := left; k <= right; k++ {
				tmp := nms[left-1]*nms[k]*nms[right+1] + dp[left][k-1] + dp[k+1][right]
				dp[left][right] = f(dp[left][right], tmp)
			}
		}
	}

	return dp[1][n]
}

func initArrays(n int) [][]int {
	re := make([][]int, n+2)

	for i := 0; i < len(re); i++ {
		tmp := make([]int, n+2)
		re[i] = tmp
	}
	return re
}

func add(nums []int) []int {
	re := make([]int, len(nums)+2)
	re[0], re[len(nums)+1] = 1, 1
	copy(re[1:], nums)
	return re
}
