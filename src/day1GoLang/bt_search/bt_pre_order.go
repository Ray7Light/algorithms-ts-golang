package BTSearch

func BTPreOrder(head BinaryNode) []int {
	var path []int

	walk(&head, &path)

	return path
}

func walk(curr *BinaryNode, path *[]int) {
	if curr == nil {
		return
	}

	*path = append(*path, curr.value)
	walk(curr.left, path)
	walk(curr.right, path)
}
