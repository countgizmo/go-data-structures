package autocomplete

type Node struct {
	children map[rune]*Node
	isWord   bool
}

type Trie struct {
	root *Node
}

func newNode() *Node {
	node := new(Node)
	node.children = make(map[rune]*Node)
	return node
}

func (t *Trie) Init() {
	t.root = newNode()
}

func (t *Trie) Insert(word string) {
	current := t.root
	for _, char := range word {
		if current.children == nil {
			current.children = make(map[rune]*Node)
		} else if current.children[char] == nil {
			current.children[char] = newNode()
		}
		current = current.children[char]
	}
	current.isWord = true
}

func buildWords(node *Node, visited *[]rune, words *[]string) {
	if node == nil {
		return
	}

	if node.isWord {
		*words = append(*words, string(*visited))
	}

	for nextChar, nextNode := range node.children {
		temp := append(*visited, nextChar)
		buildWords(nextNode, &temp, words)
	}
}

func (t *Trie) findNodeByPrefix(prefix string) *Node {
	current := t.root

	for _, ch := range prefix {
		current = current.children[ch]
	}

	return current
}

func (t *Trie) GetWords(prefix string) []string {
	var result []string
	visited := []rune{}
	startNode := t.findNodeByPrefix(prefix)

	for _, ch := range prefix {
		visited = append(visited, ch)
	}

	buildWords(startNode, &visited, &result)
	return result
}
