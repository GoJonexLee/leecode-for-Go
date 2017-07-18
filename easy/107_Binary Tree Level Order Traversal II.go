package easy

func levelOrderBottom(root *TreeNode) [][]int {
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
    
    n := len(re)
    for i := 0; i < n/2; i++ {
        re[i], re[n-i-1] = re[n-i-1], re[i]
    }
    return re
}