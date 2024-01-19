package merkle_tree

import (
	"fmt"
	"math"
	"strings"

	"example.com/ivanj26/merkle_tree/hasher"
)

func newMerkleTree(root *MerkleNode, size uint) MerkleTree {
	return MerkleTree{
		size: size,
		root: root,
	}
}

func BuildMerkleTree(transactions []string) MerkleTree {
	length := len(transactions)
	size := uint(math.Ceil(math.Log2(float64(length))) + 1)

	nodes := make([]*MerkleNode, 0)
	for _, trx := range transactions {
		hashed := hasher.HashSHA256(trx)
		nodes = append(nodes, newMerkleNode(hashed, nil, nil))
	}

	root := makeRoot(nodes)
	return newMerkleTree(root, size)
}

func makeRoot(nodes []*MerkleNode) *MerkleNode {
	length := len(nodes)
	if length == 1 {
		return nodes[0]
	}

	list := make([]*MerkleNode, 0)
	for idx := 0; idx < length; idx += 2 {
		curr := nodes[idx]

		if idx+1 >= length {
			list = append(list, curr)
			break
		}

		next := nodes[idx+1]
		combinedHash := fmt.Sprintf("%s%s", curr.value, next.value)
		parentHash := hasher.HashSHA256(combinedHash)

		list = append(list, newMerkleNode(parentHash, curr, next))
	}

	return makeRoot(list)
}

func (t *MerkleTree) AddNode(data any) {
	newNode := newMerkleNode(hasher.HashSHA256(data), nil, nil)
	t.root = addNodeRecursively(t.root, newNode)
}

func addNodeRecursively(current *MerkleNode, newNode *MerkleNode) *MerkleNode {
	if current == nil {
		return newNode
	}

	if newNode.value.(string) < current.value.(string) {
		current.left = addNodeRecursively(current.left, newNode)
	} else {
		current.right = addNodeRecursively(current.right, newNode)
	}

	var combinedHash string
	if current.left != nil {
		combinedHash = fmt.Sprint(current.left.value)
	}

	if current.right != nil {
		combinedHash += fmt.Sprint(current.right.value)
		combinedHash = hasher.HashSHA256(combinedHash)
	}

	current.value = combinedHash
	return current
}

func (t *MerkleTree) PrettyPrint() {
	t.printNode(t.root, 0)
}

func (t *MerkleTree) printNode(node *MerkleNode, depth int) {
	if node == nil {
		return
	}

	indent := strings.Repeat("   ", depth)

	fmt.Printf("%s%s\n", indent, node.GetValue())
	t.printNode(node.left, depth+1)
	t.printNode(node.right, depth+1)
}

func (t *MerkleTree) Verify(input any) bool {
	hash := hasher.HashSHA256(input)
	sibling, isLeft := findSiblingOf(hash, t.root)

	for sibling != nil && sibling.GetValue() != t.root.GetValue() {
		var val string
		if isLeft {
			val = fmt.Sprintf("%s%s", sibling.value, hash)
		} else {
			val = fmt.Sprintf("%s%s", hash, sibling.value)
		}

		hash = hasher.HashSHA256(val)
		sibling, isLeft = findSiblingOf(hash, t.root)
	}

	isFound := sibling != nil && sibling.GetValue() == t.root.GetValue()
	return isFound
}

func findSiblingOf(hash string, node *MerkleNode) (*MerkleNode, bool) {
	if node.GetValue() == hash {
		return node, false
	}
	if node.GetRight() == nil && node.GetLeft() == nil {
		return nil, false
	}
	if node.GetLeft() != nil && node.GetLeft().GetValue() == hash {
		return node.GetRight(), false
	}
	if node.GetRight() != nil && node.GetRight().GetValue() == hash {
		return node.GetLeft(), true
	}

	sibling, isLeft := findSiblingOf(hash, node.GetLeft())
	if sibling == nil {
		return findSiblingOf(hash, node.GetRight())
	}
	return sibling, isLeft
}
