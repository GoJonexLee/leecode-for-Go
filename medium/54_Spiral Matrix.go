package medium

func spiralOrder(matrix [][]int) []int {
    re := []int{}
    
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return re
    }
    
    m, n := len(matrix), len(matrix[0])
    var num, level int
    if m > n {
        num = n
    } else {
        num = m
    }
    
    c := num /2
    for level = 0; level < c; level++ {
        for i := level; i < n - level - 1; i++ {
            re = append(re, matrix[level][i])
        }
        for i := level; i < m - level - 1; i++ {
            re = append(re, matrix[i][n-level-1])
        }
        for i := n - level - 1; i > level; i-- {
            re = append(re, matrix[m-level-1][i])
        }
        for i := m - level - 1; i > level; i-- {
            re = append(re, matrix[i][level])
        }
    }
    
    if num % 2 != 0 {
        if m <= n {
            for i := level; i < n - level; i++ {
                re = append(re, matrix[level][i])
            }
        } else {
            for i := level; i < m - level; i++ {
                re = append(re, matrix[i][level])
            }
        }
    }
    
    return re
}