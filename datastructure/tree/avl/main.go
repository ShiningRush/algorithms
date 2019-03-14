package main

import (
	"fmt"
)

type BinarySearchTree struct {
	root *BinarySearchNode
}

func (t *BinarySearchTree) Insert(data int) {
	new := &BinarySearchNode{data: data}
	if t.root == nil {
		t.root = new
	} else {
		t.root.Insert(new, nil)
	}
}

func (t *BinarySearchTree) Print() {
	t.root.Print()
}

func (t *BinarySearchTree) Find(i int) bool {
	return t.root.Find(i)
}

func (t *BinarySearchTree) Remove(i int) {
	if t.root.data == i && (t.root.left == nil || t.root.right == nil) {
		if t.root.left == nil {
			t.root = t.root.right
		} else if t.root.left != nil {
			t.root = t.root.left
		} else {
			t.root = nil
		}

		return
	}

	t.root.Remove(i, nil, false)
}

type BinarySearchNode struct {
	data  int
	left  *BinarySearchNode
	right *BinarySearchNode
}

func (n *BinarySearchNode) Insert(node, pnode *BinarySearchNode) {
	if node.data > n.data {
		if n.right == nil {
			n.right = node
		} else {
			n.right.Insert(node, n)
		}
	} else {
		if n.left == nil {
			n.left = node
		} else {
			n.left.Insert(node, n)
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

func (n *BinarySearchNode) Remove(i int, pPointer *BinarySearchNode, isLeft bool) {
	if n.data == i {
		n.removeSelf(pPointer, isLeft)
	}

	if i < n.data && n.left != nil {
		n.left.Remove(i, n, true)
	}

	if i > n.data && n.right != nil {
		n.right.Remove(i, n, false)
	}
}

func (n *BinarySearchNode) removeSelf(pPointer *BinarySearchNode, isLeft bool) {
	// if node is leaf node
	if n.left == nil && n.right == nil {
		if isLeft {
			pPointer.left = nil
		} else {
			pPointer.right = nil
		}
	}

	if n.left != nil && n.right == nil {
		if isLeft {
			pPointer.left = n.left
		} else {
			pPointer.right = n.left
		}
	}

	if n.left == nil && n.right != nil {
		if isLeft {
			pPointer.left = n.right
		} else {
			pPointer.right = n.right
		}
	}

	if n.left != nil && n.right != nil {
		pr := n
		mr := n.left
		for mr.right != nil {
			pr = mr
			mr = mr.right
		}

		n.data = mr.data
		mr.removeSelf(pr, false)
	}
}

func rotation(root *BinarySearchNode, isLeft bool) *BinarySearchNode {

	// left rotation
	if isLeft {
		if root.right == nil {
			return root
		}

		pivot := root.right

		root.right = pivot.left
		pivot.left = root
		return pivot
	}

	// right rotation
	if root.left == nil {
		return root
	}

	pivot := root.left
	root.left = pivot.right
	pivot.right = root
	return pivot
}

func checkBalanceStatus(node *BinarySearchNode) BalanceStatus {
	if node.left == nil && node.right == nil {
		return Balance
	}

	hasLeftChild, hasRightChild := node.left != nil, node.right != nil
	lfChildHasLeft, lfChildHasRight, rtChildHasLeft, rtChildHasRight := false, false, false, false
	if node.left != nil {
		lfChildHasLeft, lfChildHasRight = node.left.left != nil, node.left.right != nil
	}

	if node.right != nil {
		rtChildHasLeft, rtChildHasRight = node.right.left != nil, node.right.right != nil
	}

	if hasLeftChild && lfChildHasLeft && !hasRightChild {
		return LL
	}

	if hasLeftChild && lfChildHasRight && !hasRightChild {
		return LR
	}

	if hasRightChild && rtChildHasRight && !hasLeftChild {
		return RR
	}

	if hasRightChild && rtChildHasLeft && !hasLeftChild {
		return RL
	}

	return Balance
}

type BalanceStatus uint

const (
	Balance BalanceStatus = iota
	LL
	RR
	LR
	RL
)

func main() {
	t := &BinarySearchTree{}
	ar := []int{55, 23, 11, 5, 6, 66, 22, 13, 45, 99}
	// ar := []int{55, 23, 11, 5, 6, 22, 13, 45}
	for _, v := range ar {
		t.Insert(v)
	}

	t.Print()
	fmt.Println(t.Find(22))
	t.Remove(55)
	t.Print()
}
