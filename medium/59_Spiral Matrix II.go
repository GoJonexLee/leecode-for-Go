package medium

func generateMatrix(n int) [][]int {
    re := make([][]int, n)
    for i := range re {
        re[i] = make([]int, n)
    }
    
    tmp := []int{}
    for i := 1; i <= n*n; i++ {
        tmp = append(tmp, i)
    }
    
    c, index := n / 2, 0
    for level := 0; level < c; level++ {
        for i := level; i < n - 1 - level; i++ {
            re[level][i] = tmp[index]
            index++
        }
        for i := level; i < n - level - 1; i++ {
            re[i][n-level-1] = tmp[index]
            index++
        }
        for i := n - level - 1; i > level; i-- {
            re[n-level-1][i] = tmp[index]
            index++
        }
        for i := n - level - 1; i > level; i-- {
            re[i][level] = tmp[index]
            index++
        }
    }
    
    if n % 2 != 0 {
        re[n/2][n/2] = tmp[index]
    }
    return re
}