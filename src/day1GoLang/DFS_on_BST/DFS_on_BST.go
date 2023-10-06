package DFSonBST

func DFSonBST(head BinaryNode, needle int) bool {
	return search(&head, needle)
}

func search(node *BinaryNode, needle int) bool {
	if node == nil {
		return false
	}

	if node.value == needle {
		return true
	}

	if node.value < needle {
		return 	search(node.right, needle)
	}

	return search(node.left, needle)
}
