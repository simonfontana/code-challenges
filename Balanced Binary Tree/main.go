package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
    _, balanced := depth(root)
    return balanced
}

func depth(node *TreeNode) (int, bool) {
    if node == nil {
        return 0, true
    }

	lDepth, lBalance := depth(node.Left)
	rDepth, rBalance := depth(node.Right)
	highestDepth := lDepth
	if rDepth > lDepth {
		highestDepth = rDepth
	}

	if !lBalance || !rBalance {
		return highestDepth+1, false
	}

	if lDepth-rDepth > 1 || rDepth-lDepth > 1 {
		return highestDepth+1, false
	}

	return highestDepth+1, true
}
