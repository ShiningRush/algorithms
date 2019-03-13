package main

import "fmt"

type BinarySearchTree struct {
	root *BinarySearchNode
}

func (t *BinarySearchTree) Insert(data int) {
	new := &BinarySearchNode{data: data}
	if t.root == nil {
		t.root = new
	} else {
		t.root.Insert(new)
	}
}

func (t *BinarySearchTree) Print() {
	t.root.Print()
}

func (t *BinarySearchTree) Find(i int) bool {
	return t.root.Find(i)
}

func (t *BinarySearchTree) Remove(i int) {
	t.root.Remove(i)
}

type BinarySearchNode struct {
	data  int
	left  *BinarySearchNode
	right *BinarySearchNode
}

func (n *BinarySearchNode) Insert(node *BinarySearchNode) {
	if node.data > n.data {
		if n.right == nil {
			n.right = node
		} else {
			n.right.Insert(node)
		}
	} else {
		if n.left == nil {
			n.left = node
		} else {
			n.left.Insert(node)
		}
	}
}

func (n *BinarySearchNode) Print() {
	if n.left != nil {
		n.left.Print()
	}

	fmt.Print(n.data)
	fmt.Print(" ")

	if n.right != nil {
		n.right.Print()
	}
}

func (n *BinarySearchNode) Find(i int) bool {
	if i == n.data {
		return true
	}

	if i < n.data && n.left != nil {
		return n.left.Find(i)
	}

	if i > n.data && n.right != nil {
		return n.right.Find(i)
	}

	return false
}

func (n *BinarySearchNode) Remove(i int) {
	if n.data == i {
		n.removeSelf()
	}

	if i < n.data && n.left != nil {
		n.left.Remove(i)
	}

	if i > n.data && n.right != nil {
		n.right.Remove(i)
	}
}

func (n *BinarySearchNode) removeSelf() {
	// if node is leaf node
	if n.
}

func main() {
	t := &BinarySearchTree{}
	ar := []int{55, 23, 11, 5, 6, 66, 22, 13, 45, 99}
	for _, v := range ar {
		t.Insert(v)
	}

	t.Print()
	fmt.Println(t.Find(22))
	fmt.Println(t.Find(12))
}
