package autocomplete

import (
	"testing"
)

func TestInitTrie(t *testing.T) {
	trie := new(Trie)
	trie.Init()
	if trie.root.children == nil {
		t.Error("Expected children to be present")
	}
}

func TestInsert(t *testing.T) {
	trie := new(Trie)
	trie.Init()
	trie.Insert("buy")
	trie.Insert("bun")

	if trie.root.children['b'] == nil {
		t.Error("Expected 'b' to be present")
	}
	if trie.root.children['b'].children['u'] == nil {
		t.Error("Expected 'u' to be present")
	}
	if trie.root.children['b'].children['u'].children['y'] == nil {
		t.Error("Expected 'y' to be present")
	}
	if trie.root.children['b'].children['u'].children['n'] == nil {
		t.Error("Expected 'n' to be present")
	}
	if !trie.root.children['b'].children['u'].children['y'].isWord {
		t.Error("Expected 'buy' to be a word")
	}
	if !trie.root.children['b'].children['u'].children['n'].isWord {
		t.Error("Expected 'bun' to be a word")
	}
}

func TestGetWords(t *testing.T) {
	trie := new(Trie)
	trie.Init()
	trie.Insert("buy")
	trie.Insert("bun")
	trie.Insert("bunny")

	result := trie.GetWords("b")
	if result[0] != "buy" {
		t.Errorf("Expected 'buy' got '%s'", result[0])
	}

	if result[1] != "bun" {
		t.Errorf("Expected 'bun' got '%s'", result[1])
	}

	if result[2] != "bunny" {
		t.Errorf("Expected 'bunny' got '%s'", result[2])
	}
}
