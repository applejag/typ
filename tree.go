// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import (
	"constraints"
	"fmt"
)

// Tree is an AVL tree (Adelson-Velsky and Landis tree), a type of
// self-balancing binary search tree (BST). This guarantees O(log n) operations
// on insertion, selection, and deletion.
type Tree[T constraints.Ordered] struct {
	root  *avlNode[T]
	count int
}

// Len returns the number of nodes in this tree.
func (n Tree[T]) Len() int {
	return n.count
}

func (n Tree[T]) Contains(value T) bool {
	if n.root == nil {
		return false
	}
	return n.root.contains(value)
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

func (n avlNode[T]) String() string {
	return fmt.Sprint(n.value)
}

func (n avlNode[T]) contains(value T) bool {
	// TODO: Change to binary searching
	return n.value == value ||
		(n.left != nil && n.left.contains(value)) ||
		(n.right != nil && n.right.contains(value))
}

func (n avlNode[T]) add(value T, compare func(a, b T) int) avlNode[T] {
	if compare(value, n.value) < 0 {
		if n.left == nil {
			n.left = &avlNode[T]{
				value: value,
			}
		} else {
			newLeft := n.left.add(value, compare)
			n.left = &newLeft
		}
	} else {
		if n.right == nil {
			n.right = &avlNode[T]{
				value: value,
			}
		} else {
			newRight := n.right.add(value, compare)
			n.right = &newRight
		}
	}
	return n.rebalance()
}

func (n avlNode[T]) remove(value T) avlNode[T] {
	// TODO: implement this
	return n
}

func (n avlNode[T]) rebalance() avlNode[T] {
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

func (n avlNode[T]) balance() balanceFactor {
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

func (n avlNode[T]) leftHeight() int {
	if n.left == nil {
		return 0
	}
	return n.left.height
}

func (n avlNode[T]) rightHeight() int {
	if n.right == nil {
		return 0
	}
	return n.right.height
}

func (n avlNode[T]) calcHeight() int {
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

func (n avlNode[T]) rotateLeft() avlNode[T] {
	newRoot := *n.right
	n.right = newRoot.left
	if n.right != nil {
		n.right.height = n.right.calcHeight()
	}
	n.height = n.calcHeight()
	newRoot.left = &n
	newRoot.height = newRoot.calcHeight()
	return newRoot
}

func (n avlNode[T]) rotateRight() avlNode[T] {
	newRoot := *n.left
	n.left = newRoot.right
	if n.left != nil {
		n.left.height = n.left.calcHeight()
	}
	n.height = n.calcHeight()
	newRoot.right = &n
	newRoot.height = newRoot.calcHeight()
	return newRoot
}

func (n avlNode[T]) rotateLeftRight() avlNode[T] {
	newRight := n.right.rotateRight()
	n.right = &newRight
	return n.rotateLeft()
}

func (n avlNode[T]) rotateRightLeft() avlNode[T] {
	newLeft := n.left.rotateLeft()
	n.left = &newLeft
	return n.rotateRight()
}
