package merkle_tree

type MerkleTree struct {
	root *MerkleNode
	size uint
}

type MerkleNode struct {
	left  *MerkleNode
	right *MerkleNode
	value any // Value in hash
}
