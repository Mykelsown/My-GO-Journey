package student

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return rplc
	}

	if node.Parent == nil {
		return rplc
	}

	if node.Parent.Left == node {
		node.Parent.Left = rplc
	} else if node.Parent.Right == node {
		node.Parent.Right = rplc
	}

	if rplc != nil {
		rplc.Parent = node.Parent
	}

	return root
}
