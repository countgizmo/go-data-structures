package autocomplete

import (
	"sort"
	"testing"
)

func TestInitAutocomp(t *testing.T) {
	autocomp := new(AutoCompleter)
	autocomp.Init()
	if autocomp.trie == nil {
		t.Error("Expected autocomp to have a trie")
	}
}

func TestGetSuggestions(t *testing.T) {
	autocomp := new(AutoCompleter)
	autocomp.Init()
	result := autocomp.GetSuggestions("app")
	sort.Strings(result)

	if result[0] != "applause" {
		t.Errorf("Expected 'applause' got '%s'", result[0])
	}

	if result[1] != "apple" {
		t.Errorf("Expected 'apple' got '%s'", result[1])
	}
}
