package easy

/*
问题描述： 计算两个整数之间的海明距离，即两个整数之间不同位的个数、A

算法描述：
1.亦或操作：将两个整数中位不同的置为1；
2.(tmp-1)*tmp: 将tmp中最后一个为1的位置为0
 */
func hammingDistance(x int, y int) int {
	tmp, re := x ^ y, 0
	for tmp != 0 {
		re++
		tmp = (tmp - 1) & tmp
	}
	return re
}
