package DFSonBST 

type BinaryNode struct {
	value int
	right *BinaryNode
	left  *BinaryNode
}

var Tree = BinaryNode{
	value: 20,
	right: &BinaryNode{
		value: 50,
		right: &BinaryNode{
			value: 100,
			right: nil,
			left:  nil,
		},
		left: &BinaryNode{
			value: 30,
			right: &BinaryNode{
				value: 45,
				right: nil,
				left:  nil,
			},
			left: &BinaryNode{
				value: 29,
				right: nil,
				left:  nil,
			},
		},
	},
	left: &BinaryNode{
		value: 10,
		right: &BinaryNode{
			value: 15,
			right: nil,
			left:  nil,
		},
		left: &BinaryNode{
			value: 5,
			right: &BinaryNode{
				value: 7,
				right: nil,
				left:  nil,
			},
			left: nil,
		},
	},
}

var Tree2 = BinaryNode{
	value: 20,
	right: &BinaryNode{
		value: 50,
		right: nil,
		left: &BinaryNode{
			value: 30,
			right: &BinaryNode{
				value: 45,
				right: &BinaryNode{
					value: 49,
					left:  nil,
					right: nil,
				},
				left: nil,
			},
			left: &BinaryNode{
				value: 29,
				right: nil,
				left: &BinaryNode{
					value: 21,
					right: nil,
					left:  nil,
				},
			},
		},
	},
	left: &BinaryNode{
		value: 10,
		right: &BinaryNode{
			value: 15,
			right: nil,
			left:  nil,
		},
		left: &BinaryNode{
			value: 5,
			right: &BinaryNode{
				value: 7,
				right: nil,
				left:  nil,
			},
			left: nil,
		},
	},
}
