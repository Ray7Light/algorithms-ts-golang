package BTSearch

func BTInOrder(head BinaryNode) []int {
	var path []int

	walkInOrder(&head, &path)

	return path
}

func walkInOrder(curr *BinaryNode, path *[]int) {
	if curr == nil {
		return
	}

	walkInOrder(curr.left, path)
	*path = append(*path, curr.value)
	walkInOrder(curr.right, path)
}
