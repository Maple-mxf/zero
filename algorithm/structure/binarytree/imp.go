package binarytree

import "fmt"

type node struct {
	Data  int32
	Left  *node
	Right *node
}

type Tree struct {
	Root *node
}

func NewTree() *Tree {
	t := new(Tree)
	n6 := &node{Data: 6}
	n5 := &node{Data: 5}
	n4 := &node{Data: 4}
	n3 := &node{Data: 3}
	n2 := &node{Data: 2, Left: n5, Right: n6}
	n1 := &node{Data: 1, Left: n3, Right: n4}
	n0 := &node{Data: 0, Left: n1, Right: n2}
	t.Root = n0
	return t
}

func (t *Tree) LeftOrder() {
	t.LeftOrder0(t.Root)
}
func (t *Tree) LeftOrder0(n *node) {
	if n == nil {
		return
	}
	t.LeftOrder0(n.Left)
	fmt.Printf("%d\t", n.Data)
	t.LeftOrder0(n.Right)
}

func (t *Tree) RightOrder() {
	t.RightOrder0(t.Root)
}
func (t *Tree) RightOrder0(n *node) {
	if n == nil {
		return
	}
	t.RightOrder0(n.Right)
	fmt.Printf("%d\t", n.Data)
	t.RightOrder0(n.Left)
}

func (t *Tree) LevelOrder() {
	if t.Root == nil {
		return
	}
	var queue []*node
	queue = append(queue, t.Root)
	for len(queue) > 0 {
		var ql = len(queue)
		for i := 0; i < ql; i++ {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			fmt.Printf("%d\t", queue[0].Data)
			queue = queue[1:]
		}
		fmt.Println()
	}
}
