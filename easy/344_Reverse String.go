package easy

/*
问题描述：反转字符串

算法描述：
go的字符串不能更改，因此需要根据已知字符串生成字节切片，
其次反转字节切片，最终返回该切片的字符串表示
 */

func reverseString(s string) string {
	re, n := []byte(s), len(s)
	for i := 0; i < n/2; i++ {
		re[i], re[n-i-1] = re[n-i-1], re[i]
	}
	return string(re)
}
