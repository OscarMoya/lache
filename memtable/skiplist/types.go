package skiplist

import "github.com/oscarmoya/lache/memtable"

// node represents a node in the skip list.
// It contains a key, a value, and a slice of pointers to the next nodes at different levels.
type node struct {
	key   memtable.Key
	value []byte
	next  []*node
}

type SkipList struct {
	head        *node
	level       int
	size        int
	maxLevel    int
	probability float64 // Probability for random level generation, default is 0.5
}

// New creates a new skip list with the specified maximum level.
func New(maxLevel int, probability float64) *SkipList {

	return &SkipList{
		head:        &node{next: make([]*node, maxLevel)},
		level:       0,
		size:        0,
		probability: probability,
		maxLevel:    maxLevel,
	}
}
