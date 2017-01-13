package easy

/*
问题描述：
判断给定整数是否是2的幂次方

算法描述：时间复杂度O(1), 空间复杂度O(1)
1.二的幂次方的位表示，除了最高位为1之外，其余位全为零；
2.位操作。
 */

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return (n-1)&n == 0
}