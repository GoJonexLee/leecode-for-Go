package medium

func levelOrder(root *TreeNode) [][]int {
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
        level++
        f(node.Left, level)
        f(node.Right, level)
    }
    
    f(root, 0)
    return re
}