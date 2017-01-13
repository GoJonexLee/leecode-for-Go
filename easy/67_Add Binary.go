package easy

/*
问题描述：
字符串a、b均表示二进制字符串，求a与b的和，并以字符串的形式返回

算法描述：时间复杂度O(n),空间复杂度O(n)
1.btoi记录字节对应的整数，itob记录整数对应的字节；
2.从未到头同时遍历字符串a、b，求和，并且记录进位值；
3.如果a、b不同长度，则需要计算剩余较长的字符串；
4.反转求和切片，并以字符串形式返回。
 */

func addBinary(a string, b string) string {
	btoi := map[byte]int{'0':0, '1':1}
	itob := map[int]byte{0:'0', 1:'1'}

	nt, sum := 0, 0
	i, j := len(a)-1, len(b)-1
	re := []byte{}

	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		sum = nt + btoi[a[i]] + btoi[b[j]]
		nt = sum / 2
		now := sum % 2
		re = append(re, itob[now])
	}

	for j >= 0 {
		sum = nt + btoi[b[j]]
		nt = sum / 2
		now := sum % 2
		re = append(re, itob[now])
		j--
	}

	for i >= 0 {
		sum = nt + btoi[a[i]]
		nt = sum / 2
		now := sum % 2
		re = append(re, itob[now])
		i--
	}

	if nt > 0 {
		re = append(re, itob[nt])
	}

	for i := 0; i < len(re)/2; i++ {
		re[i], re[len(re)-1-i] = re[len(re)-1-i], re[i]
	}

	return string(re)
}
