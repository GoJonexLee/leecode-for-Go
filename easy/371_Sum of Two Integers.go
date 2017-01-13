package easy

func getSum(a int, b int) int {
	re := a ^ b

	if a & b != 0 {
		return getSum(a & b << 1, a ^ b)
	}

	return re
}