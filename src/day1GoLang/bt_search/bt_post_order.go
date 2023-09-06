package BTSearch

func BTPostOrder(head BinaryNode) []int {
	var path []int

	walkPostOrder(&head, &path)

	return path
}

func walkPostOrder(curr *BinaryNode, path *[]int) {
	if curr == nil {
		return
	}

	walkPostOrder(curr.left, path)
	walkPostOrder(curr.right, path)
	*path = append(*path, curr.value)
}
