package easy

func sortedArrayToBST(nums []int) *TreeNode {
    
    var f func(int, int) *TreeNode
    f = func(low, high int) *TreeNode {
        if low > high {
            return nil
        }
        
        mid := (low + high) / 2
        node := &TreeNode{Val: nums[mid]}
        node.Left = f(low, mid-1)
        node.Right = f(mid+1, high)
        
        return node
    }
    
    return f(0, len(nums)-1)
}