package easy

import "strconv"

/*
问题描述：
3的倍数：Fizz， 5的倍数：Buzz， 6的倍数：Fizz， 15的倍数：FizzBuzz

算法表述：
倍数优先级：15 > 3 > 5 > 6 > 整数
 */

func fizzBuzz(n int) []string {
	re := []string{}
	for i := 1; i <= n; i++ {
		tmp := ""
		if i % 15 == 0 {
			tmp = "FizzBuzz"
		} else if i % 3 == 0 {
			tmp = "Fizz"
		} else if i % 5 == 0 {
			tmp = "Buzz"
		} else if i % 6 == 0 {
			tmp = "Fizz"
		}  else {
			tmp = strconv.Itoa(i)
		}
		re = append(re, tmp)
	}
	return re
}