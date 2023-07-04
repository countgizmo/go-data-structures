package bst

import (
	"testing"
)

func GetTree() Tree {
	tree := new(Tree)
	tree.Insert(12)
	tree.Insert(5)
	tree.Insert(18)
	tree.Insert(2)
	tree.Insert(9)
	tree.Insert(15)
	tree.Insert(19)
	tree.Insert(13)
	tree.Insert(17)

	return *tree
}

func TestInsert(t *testing.T) {
	tree := new(Tree)
	tree.Insert(12)
	tree.Insert(18)
	if tree.root.key != 12 {
		t.Errorf("Expected root key to be 12 got %d", tree.root.key)
	}
}

func TestMaximum(t *testing.T) {
	tree := new(Tree)

	tree.Insert(12)
	max := tree.Maximum()
	if max != 12 {
		t.Errorf("Expected max to be 12 got %d", max)
	}

	tree.Insert(18)
	max = tree.Maximum()
	if max != 18 {
		t.Errorf("Expected max to be 18 got %d", max)
	}
}

func TestMinimum(t *testing.T) {
	tree := new(Tree)
	tree.Insert(12)
	tree.Insert(18)
	min := tree.Minimum()
	if min != 12 {
		t.Errorf("Expected min to be 12 got %d", min)
	}

	tree.Insert(5)
	min = tree.Minimum()
	if min != 5 {
		t.Errorf("Expected min to be 5 got %d", min)
	}
}

func TestSearch(t *testing.T) {
	tree := GetTree()

	result := tree.Search(15)
	if result.key != 15 {
		t.Errorf("Expected to find node with key 15, got node %v", result)
	}

	result = tree.Search(777)
	if result != nil {
		t.Errorf("Expected to find nothing, got %v", result)
	}
}

func TestDelete(t *testing.T) {
	tree := GetTree()

	tree.Print()
	tree.Delete(15)
	result := tree.Search(18)
	if result.left.key != 17 {
		t.Errorf("Expected the new child to be 17, got node %v", result.left)
	}

	tree.Print()
}

func TestSuccessor(t *testing.T) {
	tree := GetTree()
	result, e := tree.Successor(12)
	if result != 13 || e != nil {
		t.Errorf("Expected successor 13 got %v", result)
	}

	result, e = tree.Successor(9)
	if result != 12 || e != nil {
		t.Errorf("Expected successor 12 got %v", result)
	}

	result, e = tree.Successor(2)
	if result != 5 || e != nil {
		t.Errorf("Expected successor 5 got %v", result)
	}
}

func TestPredecessor(t *testing.T) {
	tree := GetTree()
	result := tree.Predecessor(12)
	if result != 9 {
		t.Errorf("Expected predecessor 9 got %v", result)
	}

	result = tree.Predecessor(13)
	if result != 12 {
		t.Errorf("Expected predecessor 12 got %v", result)
	}

	result = tree.Predecessor(19)
	if result != 18 {
		t.Errorf("Expected predecessor 18 got %v", result)
	}
}
