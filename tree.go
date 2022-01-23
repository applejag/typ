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

// Clone will return a copy of this tree, with a new set of nodes. The values
// are copied as-is, so no pointers inside your value type gets a deep clone.
func (n *OrderedTree[T]) Clone() OrderedTree[T] {
	var clone OrderedTree[T]
	n.WalkPreOrder(clone.Add)
	return clone
}

// Len returns the number of nodes in this tree.
func (n *OrderedTree[T]) Len() int {
	return n.count
}

// Contains checks if a value exists in this tree by iterating the binary
// search tree.
func (n *OrderedTree[T]) Contains(value T) bool {
	if n.root == nil {
		return false
	}
	return n.root.contains(value, compare[T])
}

// Add will add another value to this tree. Duplicate values are allowed and
// are not dismissed.
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

// Remove will try to remove the first occurrence of a value from the tree.
func (n *OrderedTree[T]) Remove(value T) bool {
	if n.root == nil {
		return false
	}
	newRoot, ok := n.root.remove(value, compare[T])
	n.root = newRoot
	n.count--
	return ok
}

// Clear will reset this tree to an empty tree.
func (n *OrderedTree[T]) Clear() {
	n.root = nil
	n.count = 0
}

// WalkPreOrder will iterate all values in this tree by first visiting each
// node's value, followed by the its left branch, and then its right branch.
//
// This is useful when copying binary search trees, as inserting back in this
// order will guarantee the clone will have the exact same layout.
func (n *OrderedTree[T]) WalkPreOrder(walker func(value T)) {
	if n.root == nil {
		return
	}
	n.root.walkPreOrder(walker)
}

// WalkInOrder will iterate all values in this tree by first visiting each
// node's left branch, followed by the its own value, and then its right branch.
//
// This is useful when reading a tree's values in order, as this guarantees
// iterating them in a sorted order.
func (n *OrderedTree[T]) WalkInOrder(walker func(value T)) {
	if n.root == nil {
		return
	}
	n.root.walkInOrder(walker)
}

// WalkPostOrder will iterate all values in this tree by first visiting each
// node's left branch, followed by the its right branch, and then its own value.
//
// This is useful when deleting values from a tree, as this guarantees to always
// delete leaf nodes.
func (n *OrderedTree[T]) WalkPostOrder(walker func(value T)) {
	if n.root == nil {
		return
	}
	n.root.walkPostOrder(walker)
}

// SlicePreOrder returns a slice of values by walking the tree in pre-order.
// See WalkPreOrder for more details.
func (n *OrderedTree[T]) SlicePreOrder() []T {
	return n.slice(n.WalkPreOrder)
}

// SliceInOrder returns a slice of values by walking the tree in in-order.
// This returns all values in sorted order.
// See WalkInOrder for more details.
func (n *OrderedTree[T]) SliceInOrder() []T {
	return n.slice(n.WalkInOrder)
}

// SlicePostOrder returns a slice of values by walking the tree in post-order.
// See WalkPostOrder for more details.
func (n *OrderedTree[T]) SlicePostOrder() []T {
	return n.slice(n.WalkPostOrder)
}

func (n *OrderedTree[T]) slice(f func(f func(value T))) []T {
	slice := make([]T, 0, n.count)
	f(func(v T) {
		slice = append(slice, v)
	})
	return slice
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

func (n *avlNode[T]) walkPreOrder(f func(v T)) {
	f(n.value)
	if n.left != nil {
		n.left.walkPreOrder(f)
	}
	if n.right != nil {
		n.right.walkPreOrder(f)
	}
}

func (n *avlNode[T]) walkInOrder(f func(v T)) {
	if n.left != nil {
		n.left.walkInOrder(f)
	}
	f(n.value)
	if n.right != nil {
		n.right.walkInOrder(f)
	}
}

func (n *avlNode[T]) walkPostOrder(f func(v T)) {
	if n.left != nil {
		n.left.walkPostOrder(f)
	}
	if n.right != nil {
		n.right.walkPostOrder(f)
	}
	f(n.value)
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
		return 1 + Max(n.leftHeight(), n.rightHeight())
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
