package linkedlist

type LinkedList struct {
	Head *node
}

type node struct {
	Next *node
	Data interface{}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: &node{
			Next: nil,
		}}
}

func (imp *LinkedList) Add(n *node) {
	if imp.Head.Next == nil {
		imp.Head.Next = n
		return
	}
	p := imp.Head.Next
	for p.Next != nil {
		p = p.Next
	}
	p.Next = n
}

func (imp *LinkedList) GetNodes() (nodes []*node) {
	p := imp.Head
	for p.Next != nil {
		nodes = append(nodes, p.Next)
	}
	return
}
