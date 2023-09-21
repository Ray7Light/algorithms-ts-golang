package CompareBT

func CompareBT(tree *BinaryNode, tree2 *BinaryNode) bool {

	if tree == nil && tree2 == nil {
		return true
	}

	if tree == nil || tree2 == nil {
		return false
	}

	if tree.value != tree2.value {
		return false
	}

	return CompareBT(tree.left, tree2.left) && CompareBT(tree.right, tree2.right)
}
