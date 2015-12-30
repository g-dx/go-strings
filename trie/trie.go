package trie
import "fmt"

// Basic trie interface. Can only store strings - no associated values
type Trie interface {
	Search(key string) bool
	Insert(key string)
	Count() uint
	SizeOf() uint
	Print()
}

// Interface to support walking trie in a uniform way
type node interface {
	Children() []node
	Val() interface{}
}

// Returns a new trie with the simplest implementation
func Simple() Trie {
	return &simpleTrie{}
}

// Returns a new trie using a "left child, right sibling" implementation
func LCRS() Trie {
	return &lcrsTrie{val : rune('-')}
}

// Returns a new radix trie
func Radix() Trie {
	return &radixTrie{ edge : "-" }
}

func printTrie(n node) {
	print(n, "", true)
}

func print(n node, prefix string, isTail bool) {

	// Handle current node
	row := "├── "
	if isTail {
		row = "└── "
	}

	fmt.Printf("%v%v%v\n", prefix, row, n.Val())

	row = "|     "
	if isTail {
		row = "      "
	}
	// Print children
	children := n.Children()
	for i, child := range children {
		isTail = false;
		if i == len(children) - 1 {
			isTail = true
		}
		print(child, prefix + row, isTail)
	}
}