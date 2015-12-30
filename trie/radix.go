package trie
import (
	"unsafe"
)

// Radix tree implemented using RCLS storing only strings
// 64-bit: 40bytes + len(edge)
type radixTrie struct {
	edge  string
	child *radixTrie
	next  *radixTrie
	isWord bool
}

func (r *radixTrie) addChild(c *radixTrie) {
	if (r.child == nil) {
		r.child = c
	} else {
		r.child.addSibling(c)
	}
}

func (r *radixTrie) addSibling(c *radixTrie) {
	n := r
	for ; n.next != nil; n = n.next {
		// Skip
	}
	n.next = c
}

func (r *radixTrie) Search(s string) bool {
	cur := r
	key := s
	for {
		// Attempt to fully match an edge
		cur = cur.matchExactEdge(key)
		if cur != nil {
			// Truncate key & search remainder
			key = key[len(cur.edge):]
			if len(key) > 0 {
				continue
			}

			// Key has been totally consumed
			break
		}

		// No matching edge found
		return false
	}
	// check this node is a word
	return cur.isWord
}

func (r *radixTrie) Insert(s string) {

	cur := r
	parent := cur
	key := s
	i := 0
	for {
		// Attempt to find partially matching edge
		parent = cur
		cur, i = cur.matchPartialEdge(key)

		// No outgoing edge matches any prefix. Add new node to parent with remaining prefix.
		if cur == nil {
			parent.addChild(&radixTrie{edge : key, isWord : true})
			return
		}
		// The current node is a prefix of the key
		if i == len(cur.edge) {
			key = key[i:]
			// There is no remaining key (i.e. the original key was already present). No further work required
			if len(key) == 0 {
				return
			}
			// Process the remainder of the key
			continue
		}

		// The prefix is smaller than the current edge. We must split the current node
		// Save old child (if any). This will need fixed up. Create new child with prefix.
		oldChild := cur.child
		oldSuffix := &radixTrie{edge : cur.edge[i:], isWord:cur.isWord}

		// Correct current node to only store old prefix
		cur.edge = cur.edge[:i]
		cur.isWord = len(key[i:]) == 0 // if key has no remainder then this key was a subword of the edge we are splitting

		// Add old suffix as a child of split node
		cur.child = oldSuffix

		// Add old child of now split parent to old suffix node
		oldSuffix.child = oldChild

		// If there is a remainder add new suffix at level of old suffix
		if len(key[i:]) > 0 {
			s := string(append([]byte(nil), key[i:]...)) // Copy remaining key to prevent memory leak
			oldSuffix.addSibling(&radixTrie{edge : s, isWord:true})
		}
		return
	}
}

func (r *radixTrie) matchPartialEdge(key string) (*radixTrie, int) {
	// Has children?
	if r.child == nil {
		return nil, -1
	}
	// Check child
	i := commonPrefix(key, r.child.edge)
	if  i > 0 {
		return r.child, i
	}
	// Check siblings
	for n := r.child.next; n != nil; n = n.next {
		i = commonPrefix(key, n.edge)
		if i > 0 {
			return n, i
		}
	}
	return nil, i
}

func (r *radixTrie) matchExactEdge(key string) *radixTrie {
	n, i := r.matchPartialEdge(key)
	if n != nil && i != len(n.edge) {
		n = nil
	}
	return n
}

func commonPrefix(s1 string, s2 string) int {
	max := len(s1)
	if n := len(s2); n < max {
		max = n
	}
	i := 0
	for ; i < max; i++ {
		if s1[i] != s2[i] {
			return i
		}
	}
	return i
}

func (r *radixTrie) Count() uint {
	n := uint(1)
	if r.child != nil {
		n += r.child.Count()
	}
	if r.next != nil {
		n += r.next.Count()
	}
	return n
}

func (r *radixTrie) SizeOf() uint {
	// Calc size of struct plus slice backing storage
	size := uint(unsafe.Sizeof(r)) + uint(len(r.edge)) // Len returns no of bytes - not runes
	if r.child != nil {
		size += r.child.SizeOf()
	}
	if r.next != nil {
		size += r.next.SizeOf()
	}
	return size
}

// ===========================================================================
// Print support
// ===========================================================================

func (t *radixTrie) Print() {
	printTrie(t)
}

func (t *radixTrie) Val() interface{} {
	return t.edge
}

func (t *radixTrie) Children() []node {
	children := make([]node, 0, 3)
	if t.child != nil {
		children = append(children, t.child)
		for n := t.child.next; n != nil ; n = n.next {
			children = append(children, n)
		}
	}
	return children
}