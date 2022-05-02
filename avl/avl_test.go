// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package avl

import (
	"testing"
	"unicode/utf8"
)

type intNode = node[int]

func TestAVLNodeRotRight(t *testing.T) {
	/*
		    4
		   /
		  2
		 / \
		1   3
	*/
	tree := &intNode{
		value:  4,
		height: 2,
		left: &intNode{
			value:  2,
			height: 1,
			left: &intNode{
				value: 1,
			},
			right: &intNode{
				value: 3,
			},
		},
	}
	/*
		  2
		 / \
		1   4
		   /
		  3
	*/
	want := &intNode{
		value:  2,
		height: 2,
		left: &intNode{
			value: 1,
		},
		right: &intNode{
			value:  4,
			height: 1,
			left: &intNode{
				value: 3,
			},
		},
	}
	got := tree.rotateRight()
	assertAVLNode(t, want, got)
}

func TestAVLNodeRotRightLeft(t *testing.T) {
	/*
		  3
		 /
		1
		 \
		  2
	*/
	tree := &intNode{
		value:  3,
		height: 2,
		left: &intNode{
			value:  1,
			height: 1,
			right: &intNode{
				value: 2,
			},
		},
	}
	/*
		  2
		 / \
		1   3
	*/
	want := &intNode{
		value:  2,
		height: 1,
		left: &intNode{
			value: 1,
		},
		right: &intNode{
			value: 3,
		},
	}
	got := tree.rotateRightLeft()
	assertAVLNode(t, want, got)
}

func TestAVLNodeRotLeft(t *testing.T) {
	/*
		1
		 \
		  3
		 / \
		2   4
	*/
	tree := &intNode{
		value:  1,
		height: 2,
		right: &intNode{
			value:  3,
			height: 1,
			left: &intNode{
				value: 2,
			},
			right: &intNode{
				value: 4,
			},
		},
	}
	/*
		  3
		 / \
		1   4
		 \
		  2
	*/
	want := &intNode{
		value:  3,
		height: 2,
		left: &intNode{
			value:  1,
			height: 1,
			right: &intNode{
				value: 2,
			},
		},
		right: &intNode{
			value: 4,
		},
	}
	got := tree.rotateLeft()
	assertAVLNode(t, want, got)
}

func TestAVLNodeRotLeftRight(t *testing.T) {
	/*
		1
		 \
		  3
		 /
		2
	*/
	tree := &intNode{
		value:  1,
		height: 2,
		right: &intNode{
			value:  3,
			height: 1,
			left: &intNode{
				value: 2,
			},
		},
	}
	/*
		  2
		 / \
		1   3
	*/
	want := &intNode{
		value:  2,
		height: 1,
		left: &intNode{
			value: 1,
		},
		right: &intNode{
			value: 3,
		},
	}
	got := tree.rotateLeftRight()
	assertAVLNode(t, want, got)
}

func FuzzOrderedTree_AddRemove(f *testing.F) {
	testcases := []string{
		"abcdefg",
		"a",
		"aaaaabbbb",
	}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, str string) {
		tree := NewOrdered[rune]()
		t.Logf("using runes: %q", str)
		strLen := utf8.RuneCountInString(str)
		for _, r := range str {
			tree.Add(r)
			if !tree.Contains(r) {
				t.Errorf("just added, but contains(%q) == false", string(r))
			}
		}
		if tree.Len() != strLen {
			t.Errorf("want len=%d, got len=%d", strLen, tree.Len())
		}
		for _, r := range str {
			lenBefore := tree.Len()
			if !tree.Remove(r) {
				t.Errorf("failed to remove value %d", r)
			}
			if lenBefore-1 != tree.Len() {
				t.Errorf("len did not shrink by 1: want %d, got %d", lenBefore-1, tree.Len())
			}
		}
		if tree.Len() != 0 {
			t.Errorf("want empty, got len=%d", tree.Len())
		}
	})
}

func assertAVLNode[T comparable](t *testing.T, want, got *node[T]) {
	assertAVLNodeRec(t, want, got, "root")
}

func assertAVLNodeRec[T comparable](t *testing.T, want, got *node[T], path string) {
	if got.value != want.value {
		t.Errorf("want %[1]s.value==%[2]v, got %[1]s.value==%[3]v", path, want.value, got.value)
	}
	if got.height != want.height {
		t.Errorf("want %[1]s.height==%[2]v, got %[1]s.height==%[3]v", path, want.height, got.height)
	}
	if got.left == nil && want.left != nil {
		t.Errorf("want %[1]s.left!=nil, got %[1]s.left==nil", path)
	} else if got.left != nil && want.left == nil {
		t.Errorf("want %[1]s.left==nil, got %[1]s.left!=nil", path)
	} else if got.left != nil && want.left != nil {
		assertAVLNodeRec(t, want.left, got.left, path+".left")
	}
	if got.right == nil && want.right != nil {
		t.Errorf("want %[1]s.right!=nil, got %[1]s.right==nil", path)
	} else if got.right != nil && want.right == nil {
		t.Errorf("want %[1]s.right==nil, got %[1]s.right!=nil", path)
	} else if got.right != nil && want.right != nil {
		assertAVLNodeRec(t, want.right, got.right, path+".right")
	}
}
