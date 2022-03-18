// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package avl

import (
	"fmt"

	"gopkg.in/typ.v3"
)

// New creates a new AVL tree using a comparator function that is
// is expected to return  0 if a == b, -1 if a < b, and +1 if a > b.
func New[T comparable](compare func(a, b T) int) Tree[T] {
	return Tree[T]{
		compare: compare,
	}
}

// NewOrdered creates a new AVL tree using a default comparator function
// for any ordered type (ints, uints, floats, strings).
func NewOrdered[T typ.Ordered]() Tree[T] {
	return New(typ.Compare[T])
}

// Tree is a binary search tree (BST) for ordered Go types
// (numbers & strings), implemented as an AVL tree
// (Adelson-Velsky and Landis tree), a type of self-balancing BST. This
// guarantees O(log n) operations on insertion, searching, and deletion.
type Tree[T comparable] struct {
	compare func(a, b T) int
	root    *node[T]
	count   int
}

func (n Tree[T]) String() string {
	return fmt.Sprint(n.SliceInOrder())
}

// Clone will return a copy of this tree, with a new set of nodes. The values
// are copied as-is, so no pointers inside your value type gets a deep clone.
func (n *Tree[T]) Clone() Tree[T] {
	var clone Tree[T]
	n.WalkPreOrder(clone.Add)
	return clone
}

// Len returns the number of nodes in this tree.
func (n *Tree[T]) Len() int {
	return n.count
}

// Contains checks if a value exists in this tree by iterating the binary
// search tree.
func (n *Tree[T]) Contains(value T) bool {
	if n.root == nil {
		return false
	}
	return n.root.contains(value, n.compare)
}

// Add will add another value to this tree. Duplicate values are allowed and
// are not dismissed.
func (n *Tree[T]) Add(value T) {
	if n.root == nil {
		n.root = &node[T]{
			value: value,
		}
	} else {
		n.root = n.root.add(value, n.compare)
	}
	n.count++
}

// Remove will try to remove the first occurrence of a value from the tree.
func (n *Tree[T]) Remove(value T) bool {
	if n.root == nil {
		return false
	}
	newRoot, ok := n.root.remove(value, n.compare)
	n.root = newRoot
	n.count--
	return ok
}

// Clear will reset this tree to an empty tree.
func (n *Tree[T]) Clear() {
	n.root = nil
	n.count = 0
}

// WalkPreOrder will iterate all values in this tree by first visiting each
// node's value, followed by the its left branch, and then its right branch.
//
// This is useful when copying binary search trees, as inserting back in this
// order will guarantee the clone will have the exact same layout.
func (n *Tree[T]) WalkPreOrder(walker func(value T)) {
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
func (n *Tree[T]) WalkInOrder(walker func(value T)) {
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
func (n *Tree[T]) WalkPostOrder(walker func(value T)) {
	if n.root == nil {
		return
	}
	n.root.walkPostOrder(walker)
}

// SlicePreOrder returns a slice of values by walking the tree in pre-order.
// See WalkPreOrder for more details.
func (n *Tree[T]) SlicePreOrder() []T {
	return n.slice(n.WalkPreOrder)
}

// SliceInOrder returns a slice of values by walking the tree in in-order.
// This returns all values in sorted order.
// See WalkInOrder for more details.
func (n *Tree[T]) SliceInOrder() []T {
	return n.slice(n.WalkInOrder)
}

// SlicePostOrder returns a slice of values by walking the tree in post-order.
// See WalkPostOrder for more details.
func (n *Tree[T]) SlicePostOrder() []T {
	return n.slice(n.WalkPostOrder)
}

func (n *Tree[T]) slice(f func(f func(value T))) []T {
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

type node[T comparable] struct {
	value  T
	left   *node[T]
	right  *node[T]
	height int
}

func (n *node[T]) String() string {
	return fmt.Sprint(n.value)
}

func (n *node[T]) walkPreOrder(f func(v T)) {
	f(n.value)
	if n.left != nil {
		n.left.walkPreOrder(f)
	}
	if n.right != nil {
		n.right.walkPreOrder(f)
	}
}

func (n *node[T]) walkInOrder(f func(v T)) {
	if n.left != nil {
		n.left.walkInOrder(f)
	}
	f(n.value)
	if n.right != nil {
		n.right.walkInOrder(f)
	}
}

func (n *node[T]) walkPostOrder(f func(v T)) {
	if n.left != nil {
		n.left.walkPostOrder(f)
	}
	if n.right != nil {
		n.right.walkPostOrder(f)
	}
	f(n.value)
}

func (n *node[T]) contains(value T, compare func(a, b T) int) bool {
	return n.find(value, compare) != nil
}

func (n *node[T]) find(value T, compare func(a, b T) int) *node[T] {
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

func (n *node[T]) remove(value T, compare func(a, b T) int) (*node[T], bool) {
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

func (n *node[T]) popLeftMost() (child, leftMost *node[T]) {
	if n.left == nil {
		// Found leftmost node
		return n.right, n
	}
	newLeft, popped := n.left.popLeftMost()
	n.left = newLeft
	n.height = n.calcHeight()
	return n, popped
}

func (n *node[T]) add(value T, compare func(a, b T) int) *node[T] {
	if compare(value, n.value) < 0 {
		if n.left == nil {
			n.left = &node[T]{
				value: value,
			}
		} else {
			n.left = n.left.add(value, compare)
		}
	} else {
		if n.right == nil {
			n.right = &node[T]{
				value: value,
			}
		} else {
			n.right = n.right.add(value, compare)
		}
	}
	return n.rebalance()
}

func (n *node[T]) rebalance() *node[T] {
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

func (n *node[T]) balance() balanceFactor {
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

func (n *node[T]) leftHeight() int {
	if n.left == nil {
		return 0
	}
	return n.left.height
}

func (n *node[T]) rightHeight() int {
	if n.right == nil {
		return 0
	}
	return n.right.height
}

func (n *node[T]) calcHeight() int {
	switch {
	case n.left == nil && n.right == nil:
		return 0
	case n.left == nil:
		return 1 + n.rightHeight()
	case n.right == nil:
		return 1 + n.leftHeight()
	default:
		return 1 + typ.Max(n.leftHeight(), n.rightHeight())
	}
}

func (n *node[T]) rotateLeft() *node[T] {
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

func (n *node[T]) rotateRight() *node[T] {
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

func (n *node[T]) rotateLeftRight() *node[T] {
	n.right = n.right.rotateRight()
	return n.rotateLeft()
}

func (n *node[T]) rotateRightLeft() *node[T] {
	n.left = n.left.rotateLeft()
	return n.rotateRight()
}
