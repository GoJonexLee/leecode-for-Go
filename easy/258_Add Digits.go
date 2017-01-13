package easy

func addDigits(num int) int {
	if num < 10 {
		return num
	}

	return addDigits(num/10 + num%10)
}
