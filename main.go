package main

import (
	"fmt"

	"example.com/ivanj26/merkle_tree/merkle_tree"
)

func main() {
	m := merkle_tree.BuildMerkleTree([]string{"123", "456", "789", "012"})
	m.PrettyPrint()

	testTrx := "123"
	fmt.Printf("\nIs the transaction %s verified? %t", testTrx, m.Verify(testTrx))

	testTrx = "555"
	fmt.Printf("\nIs the transaction %s verified? %t", testTrx, m.Verify(testTrx))

	m.AddNode(testTrx)
	fmt.Println("\n\nPrint tree after added 555:")
	m.PrettyPrint()

	fmt.Printf("\nIs the transaction %s verified? %t", testTrx, m.Verify(testTrx))

	// Check if merkle tree is still valid?
	fmt.Printf("\nIs the merkle tree valid? %t", m.IsValid())
}
