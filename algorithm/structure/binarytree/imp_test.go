package binarytree

import (
	"fmt"
	"testing"
)

func TestTree_LeftOrder(t *testing.T) {
	tree := NewTree()
	tree.LeftOrder()
	fmt.Println()
	tree.RightOrder()
	fmt.Println()
	tree.LevelOrder()
}
