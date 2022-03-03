package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	// left
	// visit
	// right
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
}
