package trie
import (
	"unsafe"
)

// Background: http://courses.cs.washington.edu/courses/cse373/06sp/handouts/lecture06.pdf
// 64-bit: 24 bytes
type lcrsTrie struct {
	child  *lcrsTrie
	next   *lcrsTrie
	val    rune
	isWord bool // More memory efficient than extra "NULL" nodes
}

func (t *lcrsTrie) Search(key string) bool {
	cur := t
	outer: for _, r := range key {
		cur = cur.child
		// No child
		if cur == nil {
			return false
		}
		// Match - move to child
		if cur.val == r {
			continue
		}
		// No match - check rest of siblings
		for cur = cur.next; cur != nil; cur = cur.next {
			if cur.val == r {
				continue outer
			}
		}
		// No match
		return false
	}
	return cur.isWord
}

func (t *lcrsTrie) Insert(key string) {

	cur := t
	outer: for _, r := range key {

		// Handle nil
		if cur.child == nil {
			cur.child = &lcrsTrie{val : r}
			cur = cur.child
			continue
		}

		// Check for match
		if cur.child.val == r {
			cur = cur.child
			continue
		}

		// Check rest of siblings
		n := cur.child
		for ; n.next != nil; n = n.next {
			if n.next.val == r {
				cur = n.next
				continue outer
			}
		}

		// Insert at end
		n.next = &lcrsTrie{val : r}
		cur = n.next
	}
	cur.isWord = true
}

func (t *lcrsTrie) Count() uint {
	n := uint(1)
	if t.child != nil {
		n += t.child.Count()
	}
	if t.next != nil {
		n += t.next.Count()
	}
	return n
}

func (t *lcrsTrie) SizeOf() uint {
	return t.Count() * (uint(unsafe.Sizeof(lcrsTrie{})))
}

// ===========================================================================
// Print support
// ===========================================================================

func (t *lcrsTrie) Print() {
	printTrie(t)
}

func (t *lcrsTrie) Val() interface{} {
	return string(t.val)
}

func (t *lcrsTrie) Children() []node {
	children := make([]node, 0, 3)
	if t.child != nil {
		children = append(children, t.child)
		for n := t.child.next; n != nil ; n = n.next {
			children = append(children, n)
		}
	}
	return children
}