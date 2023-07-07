package autocomplete

var allWords = []string{
	"buy",
	"bunny",
	"bun",
	"apple",
	"airplane",
	"ape",
	"applause",
}

type AutoCompleter struct {
	trie *Trie
}

func (au *AutoCompleter) Init() {
	au.trie = new(Trie)
	au.trie.Init()
	for _, word := range allWords {
		au.trie.Insert(word)
	}
}

func (au *AutoCompleter) GetSuggestions(query string) []string {
	return au.trie.GetWords(query)
}
