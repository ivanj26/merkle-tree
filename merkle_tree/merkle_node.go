package merkle_tree

func newMerkleNode(value any, left, right *MerkleNode) *MerkleNode {
	return &MerkleNode{
		value: value,
		left:  left,
		right: right,
	}
}

func (n *MerkleNode) SetLeft(node *MerkleNode) {
	n.left = node
}

func (n *MerkleNode) SetRight(node *MerkleNode) {
	n.right = node
}

func (n *MerkleNode) GetRight() *MerkleNode {
	return n.right
}

func (n *MerkleNode) GetLeft() *MerkleNode {
	return n.left
}

func (n *MerkleNode) GetValue() any {
	return n.value
}
