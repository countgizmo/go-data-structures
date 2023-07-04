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

func min(node *Node) *Node {
	current := node
	for current.left != nil {
		current = current.left
	}

	return current
}

func (t *Tree) Minimum() int {
	return min(t.root).key
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
		return min(node.right).key, nil
	}

	current := node
	currentParent := node.parent

	for currentParent != nil && current == currentParent.right {
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

	for currentParent != nil && current == currentParent.left {
		current = currentParent
		currentParent = currentParent.parent
	}

	return currentParent.key
}

func (t *Tree) transplant(dest *Node, src *Node) {
	if dest.parent == nil {
		t.root = src
	} else if dest == dest.parent.left {
		dest.parent.left = src
	} else {
		dest.parent.right = src
	}

	if src != nil {
		src.parent = dest.parent
	}
}

func (t *Tree) Delete(target int) {
	node := t.Search(target)

	if node.left == nil {
		t.transplant(node, node.right)
	} else if node.right == nil {
		t.transplant(node, node.left)
	} else {
		successor := min(node.right)
		if successor.parent != node {
			t.transplant(successor, successor.right)
			successor.right = node.right
			successor.right.parent = successor
		}
		t.transplant(node, successor)
		successor.left = node.left
		successor.left.parent = successor
	}
}

func (t *Tree) Print() {
	var nodes []*Node
	nodes = append(nodes, t.root)
	nodes = append(nodes, nil)

	var current *Node
	var currentIdx = 0

	for len(nodes) > currentIdx {
		current = nodes[currentIdx]
		currentIdx++

		if current == nil {
			fmt.Println()
			if currentIdx == len(nodes) {
				break
			}
			nodes = append(nodes, nil)
		} else {
			fmt.Print(current.key, "\t")
			if current.left != nil {
				nodes = append(nodes, current.left)
			}

			if current.right != nil {
				nodes = append(nodes, current.right)
			}
		}
	}
}
