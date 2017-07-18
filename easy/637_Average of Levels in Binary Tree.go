package easy

func averageOfLevels(root *TreeNode) []float64 {
    re := [][]int{}
    
    var f func(*TreeNode, int)
    f = func(node *TreeNode, level int) {
        if node == nil {
            return
        }
        
        if len(re) == level {
            re = append(re, []int{})
        }
        re[level] = append(re[level], node.Val)
        f(node.Left, level+1)
        f(node.Right, level+1)
    }
    f(root, 0)
    
    sum := func(nums []int) float64 {
        var re float64
        for _, num := range nums {
            re += float64(num)
        }
        
        return re / float64(len(nums))
    }
    
    result := make([]float64, len(re))
    for i := 0; i < len(result); i++ {
        result[i] = sum(re[i])
    }
    
    return result
}