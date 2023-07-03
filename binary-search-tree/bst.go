package bst

import (
	"fmt"
)

type Node struct {
	parent, left, right *Node
	key                 int
}

type Tree struct {
	root *Node
}

type treeError struct {
	arg     int
	message string
}

func (e *treeError) Error() string {
	return fmt.Sprintf("%s : %d", e.message, e.arg)
}

func (t *Tree) init(rootKey int) {
	if t.root == nil {
		t.root = new(Node)
	}
	t.root.key = rootKey
}

func (t *Tree) Insert(newKey int) {
	var newParent *Node

	for current := t.root; current != nil; {
		newParent = current
		if newKey < current.key {
			current = current.left
		} else {
			current = current.right
		}
	}

	if newParent == nil {
		t.init(newKey)
	} else if newParent.key > newKey {
		newParent.left = &Node{key: newKey, parent: newParent}
	} else {
		newParent.right = &Node{key: newKey, parent: newParent}
	}
}

func max(node *Node) int {
	current := node
	for current.right != nil {
		current = current.right
	}

	return current.key
}

func (t *Tree) Maximum() int {
	return max(t.root)
}

func min(node *Node) int {
	current := node
	for current.left != nil {
		current = current.left
	}

	return current.key
}

func (t *Tree) Minimum() int {
	return min(t.root)
}

func (t *Tree) Search(needle int) *Node {
	current := t.root
	for current != nil && current.key != needle {
		if needle < current.key {
			current = current.left
		} else {
			current = current.right
		}
	}

	return current
}

func (t *Tree) Successor(target int) (int, error) {
	node := t.Search(target)

	if node.right != nil {
		return min(node.right), nil
	}

	current := node
	currentParent := node.parent

	for currentParent != nil && current.key == currentParent.right.key {
		current = currentParent
		currentParent = currentParent.parent
	}

	if currentParent != nil {
		return currentParent.key, nil
	}

	return -1, &treeError{target, "Couldn't find successor"}
}

func (t *Tree) Predecessor(target int) int {
	node := t.Search(target)

	if node.left != nil {
		return max(node.left)
	}

	current := node
	currentParent := current.parent

	for currentParent != nil && current.key == currentParent.left.key {
		current = currentParent
		currentParent = currentParent.parent
	}

	return currentParent.key
}
