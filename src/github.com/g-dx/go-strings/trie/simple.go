package trie
import (
	"fmt"
	"unsafe"
)

const alphabetSize = 30

type simpleTrie struct {
	alphabet [alphabetSize]*simpleTrie
	isWord   bool
}

func (t *simpleTrie) Search(key string) bool {
	cur := t
	for _, r := range key {
		if cur = cur.alphabet[toIndex(r)]; cur == nil {
			return false
		}
	}
	return cur.isWord
}

func (t *simpleTrie) Insert(key string) {
	cur := t
	for _, r := range key {
		child := cur.alphabet[toIndex(r)]
		if child == nil {
			n := &simpleTrie{}
			cur.alphabet[toIndex(r)] = n
			cur = n
		} else {
			cur = child
		}
	}
	cur.isWord = true
}

func (t *simpleTrie) Count() uint {
	n := uint(1)
	for _, child := range t.alphabet {
		if child != nil {
			n += child.Count()
		}
	}
	return n
}

func (t *simpleTrie) SizeOf() uint {
	return t.Count() * (uint(unsafe.Sizeof(simpleTrie{})))
}

func toIndex(r rune) int {
	switch {
	// ASCII a-z
	case r >= 65 && r <= 90:
		return int (r - 65)
	// ASCII A-Z
	case r >= 97 && r <= 122:
		return int(r - 97)
	// Extra characters present in "dict/words.txt"
	case r == '-':
		return 26
	case r == '\'':
		return 27
	case r == '2':
		return 28
	case r == '3':
		return 29
	}
	panic(fmt.Sprintf("Rune: '%v' is not in alphabet!", string(r)))
}

func fromIndex(i int) rune {
	switch {
	// ASCII a-z
	case i >= 0 && i <= 25:
		return rune(i + 65)
	// Extra characters present in "dict/words.txt"
	case i == 26:
		return '-'
	case i == 27:
		return '\''
	case i == 28:
		return '2'
	case i == 29:
		return '3'
	}
	panic(fmt.Sprintf("Index: '%v' is not in alphabet!", i))
}

// ===========================================================================
// Print support
// ===========================================================================

func (t *simpleTrie) Print() {
	printTrie(&simpleNode{ v : "", n : t })
}

func (t *simpleNode) Val() interface{} {
	return t.v
}

type simpleNode struct {
	v string
	n *simpleTrie
}

func (t *simpleNode) Children() []node {
	children := make([]node, 0, 3)
	for i, child := range t.n.alphabet {
		if child != nil {
			children = append(children, &simpleNode{v : string(fromIndex(i)), n : child})
		}
	}
	return children
}
