// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import (
	"constraints"
	"fmt"
)

// OrderedTree is a binary search tree (BST) for ordered Go types
// (numbers & strings), implemented as an AVL tree
// (Adelson-Velsky and Landis tree), a type of self-balancing BST. This
// guarantees O(log n) operations on insertion, searching, and deletion.
type OrderedTree[T constraints.Ordered] struct {
	root  *avlNode[T]
	count int
}

// Len returns the number of nodes in this tree.
func (n *OrderedTree[T]) Len() int {
	return n.count
}

func (n *OrderedTree[T]) Contains(value T) bool {
	if n.root == nil {
		return false
	}
	return n.root.contains(value, compare[T])
}

func (n *OrderedTree[T]) Add(value T) {
	if n.root == nil {
		n.root = &avlNode[T]{
			value: value,
		}
	} else {
		n.root = n.root.add(value, compare[T])
	}
	n.count++
}

func (n *OrderedTree[T]) Remove(value T) bool {
	if n.root == nil {
		return false
	}
	newRoot, ok := n.root.remove(value, compare[T])
	n.root = newRoot
	n.count--
	return ok
}

type balanceFactor int8

const (
	balanceBalanced   balanceFactor = 0
	balanceRightHeavy balanceFactor = 1
	balanceLeftHeavy  balanceFactor = -1
)

type avlNode[T comparable] struct {
	value  T
	left   *avlNode[T]
	right  *avlNode[T]
	height int
}

func (n *avlNode[T]) String() string {
	return fmt.Sprint(n.value)
}

func (n *avlNode[T]) contains(value T, compare func(a, b T) int) bool {
	return n.find(value, compare) != nil
}

func (n *avlNode[T]) find(value T, compare func(a, b T) int) *avlNode[T] {
	current := n
	for {
		switch {
		case current.value == value:
			return current
		case current.left != nil && compare(value, current.value) < 0:
			current = current.left
		case current.right != nil:
			current = current.right
		default:
			return nil
		}
	}
}

func (n *avlNode[T]) remove(value T, compare func(a, b T) int) (*avlNode[T], bool) {
	if n.value == value {
		switch {
		case n.left == nil && n.right == nil:
			// Leaf node. No special behavior needed
			return nil, true
		case n.left == nil:
			// Single child: right
			return n.right, true
		case n.right == nil:
			// Single child: left
			return n.left, true
		default:
			// Two children
			newRight, leftMost := n.right.popLeftMost()
			leftMost.left = n.left
			leftMost.right = newRight
			leftMost.height = leftMost.calcHeight()
			return leftMost.rebalance(), true
		}
	}
	if n.left != nil && compare(value, n.value) < 0 {
		if newNode, ok := n.left.remove(value, compare); ok {
			n.left = newNode
			n.height = n.calcHeight()
			return n.rebalance(), true
		}
	} else if n.right != nil {
		if newNode, ok := n.right.remove(value, compare); ok {
			n.right = newNode
			n.height = n.calcHeight()
			return n.rebalance(), true
		}
	}
	return n, false
}

func (n *avlNode[T]) popLeftMost() (child, leftMost *avlNode[T]) {
	if n.left == nil {
		// Found leftmost node
		return n.right, n
	}
	newLeft, popped := n.left.popLeftMost()
	n.left = newLeft
	n.height = n.calcHeight()
	return n, popped
}

func (n *avlNode[T]) add(value T, compare func(a, b T) int) *avlNode[T] {
	if compare(value, n.value) < 0 {
		if n.left == nil {
			n.left = &avlNode[T]{
				value: value,
			}
		} else {
			n.left = n.left.add(value, compare)
		}
	} else {
		if n.right == nil {
			n.right = &avlNode[T]{
				value: value,
			}
		} else {
			n.right = n.right.add(value, compare)
		}
	}
	return n.rebalance()
}

func (n *avlNode[T]) rebalance() *avlNode[T] {
	if n.balance() == balanceRightHeavy {
		if n.right != nil && n.right.balance() == balanceLeftHeavy {
			return n.rotateLeftRight()
		}
		return n.rotateLeft()
	} else if n.balance() == balanceLeftHeavy {
		if n.left != nil && n.left.balance() == balanceRightHeavy {
			return n.rotateRightLeft()
		}
		return n.rotateRight()
	}
	return n
}

func (n *avlNode[T]) balance() balanceFactor {
	leftHeight := n.leftHeight()
	rightHeight := n.rightHeight()
	if leftHeight-rightHeight > 1 {
		return balanceLeftHeavy
	}
	if rightHeight-leftHeight > 1 {
		return balanceRightHeavy
	}
	return balanceBalanced
}

func (n *avlNode[T]) leftHeight() int {
	if n.left == nil {
		return 0
	}
	return n.left.height
}

func (n *avlNode[T]) rightHeight() int {
	if n.right == nil {
		return 0
	}
	return n.right.height
}

func (n *avlNode[T]) calcHeight() int {
	switch {
	case n.left == nil && n.right == nil:
		return 0
	case n.left == nil:
		return 1 + n.rightHeight()
	case n.right == nil:
		return 1 + n.leftHeight()
	default:
		return 1 + max(n.leftHeight(), n.rightHeight())
	}
}

func (n *avlNode[T]) rotateLeft() *avlNode[T] {
	prevRoot := *n
	newRoot := prevRoot.right
	prevRoot.right = newRoot.left
	if prevRoot.right != nil {
		prevRoot.right.height = prevRoot.right.calcHeight()
	}
	prevRoot.height = prevRoot.calcHeight()
	newRoot.left = &prevRoot
	newRoot.height = newRoot.calcHeight()
	return newRoot
}

func (n *avlNode[T]) rotateRight() *avlNode[T] {
	prevRoot := *n
	newRoot := prevRoot.left
	prevRoot.left = newRoot.right
	if prevRoot.left != nil {
		prevRoot.left.height = prevRoot.left.calcHeight()
	}
	prevRoot.height = prevRoot.calcHeight()
	newRoot.right = &prevRoot
	newRoot.height = newRoot.calcHeight()
	return newRoot
}

func (n *avlNode[T]) rotateLeftRight() *avlNode[T] {
	n.right = n.right.rotateRight()
	return n.rotateLeft()
}

func (n *avlNode[T]) rotateRightLeft() *avlNode[T] {
	n.left = n.left.rotateLeft()
	return n.rotateRight()
}
